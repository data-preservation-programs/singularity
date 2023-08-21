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
	name string
	fs   fs.Fs
}

func (h RCloneHandler) Name() string {
	return h.name
}

func (h RCloneHandler) Write(ctx context.Context, path string, in io.Reader) (fs.Object, error) {
	objInfo := object.NewStaticObjectInfo(path, time.Now(), -1, true, nil, nil)
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

	slices.SortFunc(entries, func(i, j fs.DirEntry) bool {
		return strings.Compare(i.Remote(), j.Remote()) < 0
	})

	startScanning := last == "" // Start scanning immediately if 'last' is empty.
	for _, entry := range entries {
		switch v := entry.(type) {
		case fs.Directory:
			dirPath := v.Remote()
			// If 'last' starts with directory path followed by a slash, scan inside the directory with the remaining path.
			if strings.HasPrefix(last, dirPath+"/") {
				err = h.scan(ctx, dirPath, last, ch)
			} else if startScanning || strings.Compare(dirPath, last) > 0 {
				// If we have started scanning or the directory is greater than 'last', scan inside without 'last' param.
				err = h.scan(ctx, dirPath, "", ch)
			}
			if err != nil {
				return errors.WithStack(err)
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

func (h RCloneHandler) Read(ctx context.Context, path string, offset int64, length int64) (io.ReadCloser, fs.Object, error) {
	logger.Debugw("Read: reading path", "type", h.fs.Name(), "root", h.fs.Root(), "path", path, "offset", offset, "length", length)
	object, err := h.fs.NewObject(ctx, path)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
	}
	if length == 0 {
		return &EmptyReadCloser{}, object, nil
	}
	option := &fs.RangeOption{Start: offset, End: offset + length - 1}
	reader, err := object.Open(ctx, option)
	return reader, object, errors.WithStack(err)
}

func NewRCloneHandler(ctx context.Context, s model.Storage) (*RCloneHandler, error) {
	_, ok := BackendMap[s.Type]
	registry, err := fs.Find(s.Type)
	if !ok || err != nil {
		return nil, errors.Wrapf(ErrBackendNotSupported, "type: %s", s.Type)
	}

	f, err := registry.NewFs(ctx, s.Type, s.Path, configmap.Simple(s.Metadata))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create RClone backend %s: %s", s.Type, s.Path)
	}

	return &RCloneHandler{s.Name, f}, nil
}
