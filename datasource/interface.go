package datasource

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/encryption"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Entry struct {
	Error        error
	ScannedAt    time.Time
	Type         model.ItemType
	Path         string
	Size         uint64
	LastModified *time.Time
}

type ReadAtCloser interface {
	io.ReaderAt
	io.Closer
}

type Handler interface {
	Scan(ctx context.Context, path string, last string) <-chan Entry
	Read(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error)
	Open(ctx context.Context, path string) (ReadAtCloser, error)
	CheckItem(ctx context.Context, path string) (uint64, *time.Time, error)
}

type HandlerResolver interface {
	GetHandler(source model.Source) (Handler, error)
}

type DefaultHandlerResolver struct {
	cache sync.Map
}

func NewDefaultHandlerResolver() *DefaultHandlerResolver {
	return &DefaultHandlerResolver{cache: sync.Map{}}
}

func (*DefaultHandlerResolver) getFileSystemHandler() Handler {
	return Filesystem{}
}

func (*DefaultHandlerResolver) getHTTPHandler(meta model.Metadata) (Handler, error) {
	metadata, err := meta.GetHTTPMetadata()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get http metadata")
	}

	site := Site{Headers: metadata.Headers}
	return site, nil
}

func (*DefaultHandlerResolver) getS3Handler(meta model.Metadata) (Handler, error) {
	metadata, err := meta.GetS3Metadata()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get s3 metadata")
	}

	secret, err := encryption.DecryptFromBase64String(metadata.SecretAccessKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decrypt secret access key")
	}
	s3, err := NewS3(context.Background(), metadata.Region, metadata.Endpoint, metadata.AccessKeyID, string(secret))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create s3 client")
	}

	return s3, nil
}

func (d *DefaultHandlerResolver) GetHandler(source model.Source) (Handler, error) {
	if handler, ok := d.cache.Load(source.ID); ok {
		//nolint: forcetypeassert
		return handler.(Handler), nil
	}
	var handler Handler
	var err error
	switch source.Type {
	case model.Dir, model.Upload:
		handler = d.getFileSystemHandler()
	case model.S3Path:
		handler, err = d.getS3Handler(source.Metadata)
	case model.Website:
		handler, err = d.getHTTPHandler(source.Metadata)
	default:
		return nil, errors.New("not supported source type: " + string(source.Type))
	}

	if err != nil {
		return nil, errors.Wrap(err, "failed to get handler")
	}

	d.cache.Store(source.ID, handler)
	return handler, nil
}

func ResolveSourceType(path string) (model.SourceType, string, error) {
	if strings.HasPrefix(path, "s3://") {
		return model.S3Path, strings.TrimRight(path, "/"), nil
	}

	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return model.Website, strings.TrimRight(path, "/"), nil
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", "", err
	}

	if !info.IsDir() {
		return "", "", errors.New("path is not a directory")
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", "", errors.Wrap(err, "failed to get absolute path")
	}

	abs = filepath.ToSlash(abs)
	if len(abs) > 1 {
		abs = strings.TrimRight(abs, "/")
	}
	return model.Dir, abs, nil
}
