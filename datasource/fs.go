package datasource

import (
	"context"
	"github.com/data-preservation-programs/go-singularity/model"
	"github.com/karrick/godirwalk"
	"github.com/pkg/errors"
	"io"
	"os"
	"strings"
	"time"
)

type Filesystem struct{}

func (Filesystem) Open(ctx context.Context, path string, offset uint64, length uint64) (io.ReadCloser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open file")
	}

	if _, err := file.Seek(int64(offset), io.SeekStart); err != nil {
		file.Close()
		return nil, errors.Wrap(err, "failed to seek to offset")
	}

	limitedReader := io.LimitReader(file, int64(length))

	// Wrap the limitedReader and file in a custom ReadCloser to handle both readers and closing the file.
	readCloser := &struct {
		io.Reader
		io.Closer
	}{
		Reader: limitedReader,
		Closer: file,
	}

	return readCloser, nil
}

func (Filesystem) CheckItem(ctx context.Context, path string) (uint64, *time.Time, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to stat file")
	}

	var lastModified *time.Time
	modTime := info.ModTime()
	if modTime != (time.Time{}) {
		lastModified = &modTime
	}

	return uint64(info.Size()), lastModified, nil
}

func (Filesystem) Scan(ctx context.Context, path string, last string) <-chan Entry {
	entryChan := make(chan Entry)
	go func() {
		defer close(entryChan)

		err := godirwalk.Walk(path, &godirwalk.Options{
			Unsorted:          false,
			AllowNonDirectory: true,
			Callback: func(currentPath string, dirent *godirwalk.Dirent) error {
				if dirent.IsDir() && !strings.HasPrefix(last, currentPath) && currentPath <= last {
					return godirwalk.SkipThis
				}
				if !dirent.IsDir() && currentPath <= last {
					return godirwalk.SkipThis
				}
				if dirent.IsDir() {
					return nil
				}

				info, err := os.Stat(currentPath)
				if err != nil {
					entryChan <- Entry{Error: err}
					return nil
				}

				var lastModified *time.Time
				modTime := info.ModTime()
				if modTime != (time.Time{}) {
					lastModified = &modTime
				}

				entry := Entry{
					ScannedAt:    time.Now(),
					Type:         model.File,
					Path:         currentPath,
					Size:         uint64(info.Size()),
					LastModified: lastModified,
				}

				select {
				case <-ctx.Done():
					return ctx.Err()
				case entryChan <- entry:
				}

				return nil
			},
			ErrorCallback: func(s string, err error) godirwalk.ErrorAction {
				entryChan <- Entry{Error: err}
				return godirwalk.SkipNode
			},
			FollowSymbolicLinks: true,
		})

		if err != nil {
			entryChan <- Entry{Error: err}
		}
	}()

	return entryChan
}
