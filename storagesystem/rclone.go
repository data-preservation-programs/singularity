package storagesystem

import (
	"context"
	"io"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
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
	name     string
	fs       fs.Fs
	fsNoHead fs.Fs
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
	logger.Debugw("About: getting usage", "type", h.fs.Name(), "root", h.fs.Root())
	if h.fs.Features().About != nil {
		return h.fs.Features().About(ctx)
	}

	return nil, errors.Wrapf(ErrGetUsageNotSupported, "backend: %s", h.fs.String())
}

func (h RCloneHandler) List(ctx context.Context, path string) ([]fs.DirEntry, error) {
	logger.Debugw("List: listing path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path)
	return h.fs.List(ctx, path)
}

func (h RCloneHandler) scan(ctx context.Context, path string, last string, ch chan<- Entry) error {
	logger.Debugw("Scan: listing path", "type", h.fs.String(), "root", h.fs.Root(), "path", path, "last", last)
	entries, err := h.fs.List(ctx, path)
	if err != nil {
		err = errors.Wrapf(err, "list path: %s", path)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case ch <- Entry{Error: err}:
		}
		return errors.WithStack(err)
	}

	slices.SortFunc(entries, func(i, j fs.DirEntry) int {
		return strings.Compare(i.Remote(), j.Remote())
	})

	startScanning := last == "" // Start scanning immediately if 'last' is empty.
	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Directory:
			dirPath := v.Remote()
			switch {
			case strings.HasPrefix(last, dirPath+"/"):
				// If 'last' starts with directory path followed by a slash, scan inside the directory with the remaining path.
				err = h.scan(ctx, dirPath, last, ch)
				if err != nil {
					return errors.WithStack(err)
				}
			case startScanning || strings.Compare(dirPath, last) > 0:
				// If we have started scanning or the directory is greater than 'last', scan inside without 'last' param.
				err = h.scan(ctx, dirPath, "", ch)
				if err != nil {
					return errors.WithStack(err)
				}
			default:
				continue
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- Entry{Dir: v}:
			}

		case fs.Object:
			// If 'last' is specified, skip entries until the first entry greater than 'last' is found.
			if !startScanning {
				if strings.Compare(entry.Remote(), last) > 0 {
					logger.Debugw("Scan: found first entry greater than last", "entry", entry.Remote(), "last", last)
					startScanning = true // Found the first entry greater than 'last', start scanning.
				} else {
					continue
				}
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case ch <- Entry{Info: v}:
			}
		}
	}

	return nil
}

func (h RCloneHandler) Scan(ctx context.Context, path string, last string) <-chan Entry {
	ch := make(chan Entry)
	go func() {
		defer close(ch)
		_ = h.scan(ctx, path, last, ch)
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

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	logger.Debugw("Read: reading path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path, "offset", offset, "length", length)
	object, err := h.fsNoHead.NewObject(ctx, path)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
	}
	option := &fs.SeekOption{Offset: offset}
	reader, err := object.Open(ctx, option)

	return readCloser{
		Reader: io.LimitReader(reader, length),
		Closer: reader,
	}, object, errors.WithStack(err)
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

	return &RCloneHandler{
		name:     s.Name,
		fs:       headFS,
		fsNoHead: noHeadFS,
	}, nil
}

func overrideConfig(config *fs.ConfigInfo, s model.Storage) {
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
}
