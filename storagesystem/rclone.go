package storagesystem

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"
	"sync"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/gammazero/workerpool"
	"github.com/ipfs/go-log/v2"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/config/configmap"
	"github.com/rclone/rclone/fs/object"
	"golang.org/x/exp/slices"
)

var logger = log.Logger("storage")

var _ Handler = &RCloneHandler{}

var ErrGetUsageNotSupported = errors.New("The backend does not support getting usage quota")
var ErrBackendNotSupported = errors.New("This backend is not supported")
var ErrMoveNotSupported = errors.New("The backend does not support moving files")

type RCloneHandler struct {
	name                    string
	fs                      fs.Fs
	fsNoHead                fs.Fs
	retryMaxCount           int
	retryDelay              time.Duration
	retryBackoff            time.Duration
	retryBackoffExponential float64
	scanConcurrency         int
}

func (h RCloneHandler) Name() string {
	return h.name
}

func (h RCloneHandler) Write(ctx context.Context, path string, in io.Reader) (fs.Object, error) {
	objInfo := object.NewStaticObjectInfo(path, time.Now(), -1, true, nil, nil)
	if strings.HasSuffix(path, ".car") {
		objInfo = objInfo.WithMimeType("application/vnd.ipfs.car")
	}
	return h.fs.Put(ctx, in, objInfo)
}

func (h RCloneHandler) Move(ctx context.Context, from fs.Object, to string) (fs.Object, error) {
	if h.fs.Features().Move != nil {
		return h.fs.Features().Move(ctx, from, to)
	}
	return nil, errors.Wrapf(ErrMoveNotSupported, "backend: %s", h.fs.String())
}

func (h RCloneHandler) Remove(ctx context.Context, obj fs.Object) error {
	return obj.Remove(ctx)
}

func (h RCloneHandler) About(ctx context.Context) (*fs.Usage, error) {
	logger.Debugw("About: getting usage", "type", h.fs.Name())
	if h.fs.Features().About != nil {
		return h.fs.Features().About(ctx)
	}

	return nil, errors.Wrapf(ErrGetUsageNotSupported, "backend: %s", h.fs.String())
}

func (h RCloneHandler) List(ctx context.Context, path string) ([]fs.DirEntry, error) {
	logger.Debugw("List: listing path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path)
	return h.fs.List(ctx, path)
}

func (h RCloneHandler) scan(ctx context.Context, path string, ch chan<- Entry, wp *workerpool.WorkerPool, wg *sync.WaitGroup) {
	if ctx.Err() != nil {
		return
	}
	logger.Infow("Scan: listing path", "type", h.fs.String(), "path", path)
	entries, err := h.fs.List(ctx, path)
	if err != nil {
		err = errors.Wrapf(err, "list path: %s", path)
		select {
		case <-ctx.Done():
			return
		case ch <- Entry{Error: err}:
		}
	}

	slices.SortFunc(entries, func(i, j fs.DirEntry) int {
		return strings.Compare(i.Remote(), j.Remote())
	})

	var subCount int
	for _, entry := range entries {
		entry := entry
		switch v := entry.(type) {
		case fs.Directory:
			select {
			case <-ctx.Done():
				return
			case ch <- Entry{Dir: v}:
			}

			subPath := v.Remote()
			wg.Add(1)
			wp.Submit(func() {
				h.scan(ctx, subPath, ch, wp, wg)
				wg.Done()
			})
			subCount++
		case fs.Object:
			select {
			case <-ctx.Done():
				return
			case ch <- Entry{Info: v}:
			}
		}
	}

	logger.Debugf("Scan: finished listing path, remaining %d paths to list", subCount)
}

func (h RCloneHandler) Scan(ctx context.Context, path string) <-chan Entry {
	ch := make(chan Entry, h.scanConcurrency)
	go func() {
		var wg sync.WaitGroup
		wp := workerpool.New(h.scanConcurrency)
		wg.Add(1)
		wp.Submit(func() {
			h.scan(ctx, path, ch, wp, &wg)
			wg.Done()
		})
		wg.Wait() // OK to wait while child scans continue adding to wg
		wp.StopWait()
		close(ch)
	}()
	return ch
}

func (h RCloneHandler) Check(ctx context.Context, path string) (fs.DirEntry, error) {
	logger.Debugw("Check: checking path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path)
	return h.fs.NewObject(ctx, path)
}

type readCloser struct {
	io.Reader
	io.Closer
}

type readerWithRetry struct {
	ctx                     context.Context
	object                  fs.Object
	reader                  io.ReadCloser
	offset                  int64
	retryDelay              time.Duration
	retryBackoff            time.Duration
	retryCountMax           int
	retryCount              int
	retryBackoffExponential float64
}

func (r *readerWithRetry) Close() error {
	if r.reader != nil {
		return r.reader.Close()
	}
	return nil
}

func (r *readerWithRetry) Read(p []byte) (int, error) {
	if r.ctx.Err() != nil {
		return 0, r.ctx.Err()
	}
	n, err := r.reader.Read(p)
	r.offset += int64(n)
	//nolint:errorlint
	if err == io.EOF || err == nil {
		return n, err
	}

	if r.retryCount >= r.retryCountMax {
		return n, err
	}

	// error is not EOF
	logger.Warnf("Read error: %s, retrying after %s", err, r.retryDelay)
	select {
	case <-r.ctx.Done():
		return n, errors.Join(err, r.ctx.Err())
	case <-time.After(r.retryDelay):
	}
	r.retryCount += 1
	r.retryDelay = time.Duration(float64(r.retryDelay) * r.retryBackoffExponential)
	r.retryDelay += r.retryBackoff
	r.reader.Close()
	var err2 error
	r.reader, err2 = r.object.Open(r.ctx, &fs.SeekOption{Offset: r.offset})
	if err2 != nil {
		return n, errors.Join(err, err2)
		    	logger.Warnf("Read error: %s, retrying after 5s", err)
    			time.Sleep(5 * time.Second) // 🔥 Add delay before retrying 🔥
	}
	return n, nil
}

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	logger.Debugw("Read: reading path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path, "offset", offset, "length", length)

	if length == 0 {
		object, err := h.fs.NewObject(ctx, path)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
		}
		return io.NopCloser(bytes.NewReader(nil)), object, nil
	}

	// Fetch object from rclone storage
	object, err := h.fsNoHead.NewObject(ctx, path)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
	}

	option := &fs.SeekOption{Offset: offset}

	// Open the object for reading
	reader, err := object.Open(ctx, option)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open object stream %s", path)
	}

	// Apply a 256MB buffer & wrap it in io.NopCloser to add a Close method
	bufferedReader := io.NopCloser(bufio.NewReaderSize(reader, 256*1024*1024)) // ✅ Wrapped with io.NopCloser

	readerWithRetry := &readerWithRetry{
		ctx:                     ctx,
		object:                  object,
		reader:                  bufferedReader, // ✅ Now implements io.ReadCloser
		offset:                  offset,
		retryDelay:              h.retryDelay,
		retryBackoff:            h.retryBackoff,
		retryCountMax:           h.retryMaxCount,
		retryBackoffExponential: h.retryBackoffExponential,
	}

	if length < 0 {
		return readerWithRetry, object, nil
	}

	return readCloser{
		Reader: io.LimitReader(readerWithRetry, length),
		Closer: readerWithRetry,
	}, object, nil
}

func NewRCloneHandler(ctx context.Context, s model.Storage) (*RCloneHandler, error) {
	_, ok := BackendMap[s.Type]
	registry, err := fs.Find(s.Type)
	if !ok || err != nil {
		return nil, errors.Wrapf(ErrBackendNotSupported, "type: %s", s.Type)
	}

	ctx, _ = fs.AddConfig(ctx)
	config := fs.GetConfig(ctx)
	overrideConfig(config, s)
        config.BufferSize = 256 * 1024 * 1024    // 256MB buffer for smoother streaming

	noHeadObjectConfig := make(map[string]string)
	headObjectConfig := make(map[string]string)
	for k, v := range s.Config {
		noHeadObjectConfig[k] = v
		headObjectConfig[k] = v
	}
	noHeadObjectConfig["no_head_object"] = "true"
	headObjectConfig["no_head_object"] = "false"

	noHeadFS, err := registry.NewFs(ctx, s.Type, s.Path, configmap.Simple(noHeadObjectConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create RClone backend %s: %s", s.Type, s.Path)
	}

	headFS, err := registry.NewFs(ctx, s.Type, s.Path, configmap.Simple(headObjectConfig))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create RClone backend %s: %s", s.Type, s.Path)
	}

	scanConcurrency := 1
	if s.ClientConfig.ScanConcurrency != nil {
		scanConcurrency = *s.ClientConfig.ScanConcurrency
	}

	handler := &RCloneHandler{
		name:                    s.Name,
		fs:                      headFS,
		fsNoHead:                noHeadFS,
		retryMaxCount:           10,
		retryDelay:              time.Second,
		retryBackoff:            time.Second,
		retryBackoffExponential: 1.0,
		scanConcurrency:         scanConcurrency,
	}

	if s.ClientConfig.RetryMaxCount != nil {
		handler.retryMaxCount = *s.ClientConfig.RetryMaxCount
	}
	if s.ClientConfig.RetryDelay != nil {
		handler.retryDelay = *s.ClientConfig.RetryDelay
	}
	if s.ClientConfig.RetryBackoff != nil {
		handler.retryBackoff = *s.ClientConfig.RetryBackoff
	}
	if s.ClientConfig.RetryBackoffExponential != nil {
		handler.retryBackoffExponential = *s.ClientConfig.RetryBackoffExponential
	}

	return handler, nil
}

func overrideConfig(config *fs.ConfigInfo, s model.Storage) {
	config.UseServerModTime = true
	if s.ClientConfig.ConnectTimeout != nil {
		config.ConnectTimeout = *s.ClientConfig.ConnectTimeout
	}
	if s.ClientConfig.Timeout != nil {
		config.Timeout = *s.ClientConfig.Timeout
	}
	if s.ClientConfig.ExpectContinueTimeout != nil {
		config.ExpectContinueTimeout = *s.ClientConfig.ExpectContinueTimeout
	}
	if s.ClientConfig.InsecureSkipVerify != nil {
		config.InsecureSkipVerify = true
	}
	if s.ClientConfig.NoGzip != nil {
		config.NoGzip = true
	}
	if s.ClientConfig.UserAgent != nil {
		config.UserAgent = *s.ClientConfig.UserAgent
	}
	if len(s.ClientConfig.CaCert) > 0 {
		config.CaCert = s.ClientConfig.CaCert
	}
	if s.ClientConfig.ClientCert != nil {
		config.ClientCert = *s.ClientConfig.ClientCert
	}
	if s.ClientConfig.ClientKey != nil {
		config.ClientKey = *s.ClientConfig.ClientKey
	}
	if len(s.ClientConfig.Headers) > 0 {
		for k, v := range s.ClientConfig.Headers {
			config.Headers = append(config.Headers, &fs.HTTPOption{
				Key:   k,
				Value: v,
			})
		}
	}
	if s.ClientConfig.DisableHTTP2 != nil {
		config.DisableHTTP2 = *s.ClientConfig.DisableHTTP2
	}
	if s.ClientConfig.DisableHTTPKeepAlives != nil {
		config.DisableHTTPKeepAlives = *s.ClientConfig.DisableHTTPKeepAlives
	}
	if s.ClientConfig.UseServerModTime != nil {
		config.UseServerModTime = *s.ClientConfig.UseServerModTime
	}
	config.Transfers = 1  // Only 1 file download at a time
        config.Checkers = 1   // Only 1 checker (avoids excessive HTTP requests)
        config.LowLevelRetries = 10
        config.MaxBacklog = 1
        config.BufferSize = 256 * 1024 * 1024  // 256MB buffer for smoother streaming
}
