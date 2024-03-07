package storagesystem

import (
	"context"
	"fmt"
	"io"
	"math/rand"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/gotidy/ptr"
	"github.com/rclone/rclone/fs"
	"github.com/rclone/rclone/fs/hash"
)

// IsSameEntry checks if a given model.File and a given fs.Object represent the same entry,
// based on their size, hash, and last modification time.
//
// Parameters:
//   - ctx: Context that allows for asynchronous task cancellation.
//   - file: A model.File instance representing a file in the model.
//   - object: An fs.Object instance representing a file or object in the filesystem.
//
// Returns:
//   - bool: A boolean indicating whether the given model.File and fs.Object are considered the same entry.
//   - string: A string providing details in case of a mismatch.
//
// The function performs the following checks:
//  1. Compares the sizes of 'file' and 'object'. If there is a mismatch, it returns false along with a
//     formatted string that shows the mismatched sizes.
//  2. Retrieves the last modified time of 'object'.
//  3. Identifies a supported hash type for the storage backend of 'object' and computes its hash value.
//  4. The hash computation is skipped if the storage backend is a local file system or does not support any hash types.
//  5. If both 'file' and 'object' have non-empty hash values and these values do not match,
//     it returns false along with a formatted string that shows the mismatched hash values.
//  6. Compares the last modified times of 'file' and 'object' at the nanosecond precision.
//     If there is a mismatch, it returns false along with a formatted string that shows the mismatched times.
//  7. If all the checks pass (sizes, hashes, and last modified times match), the function returns true,
//     indicating that 'file' and 'object' are considered to be the same entry.
//
// Note:
//   - In certain cases (e.g., failures during fetch), the last modified time might not be reliable.
//   - For local file systems, hash computation is skipped to avoid inefficient operations.
func IsSameEntry(ctx context.Context, file model.File, object fs.ObjectInfo) (bool, string) {
	if file.Size >= 0 && object.Size() >= 0 && file.Size != object.Size() {
		return false, fmt.Sprintf("size mismatch: %d != %d", file.Size, object.Size())
	}
	// last modified can be time.Now() if fetch failed so it may not be reliable.
	// This usually won't happen for most cloud provider i.e. S3
	// Because during scanning, the modified time is already fetched.
	lastModified := object.ModTime(ctx)
	hashValue, _ := GetHash(ctx, object)
	if file.Hash != "" && hashValue != "" && file.Hash != hashValue {
		return false, fmt.Sprintf("hash mismatch: %s != %s", file.Hash, hashValue)
	}
	return lastModified.UnixNano() == file.LastModifiedNano,
		fmt.Sprintf("last modified mismatch: %d != %d",
			lastModified.UnixNano(),
			file.LastModifiedNano)
}

// GetHash computes the hash value of a given object based on the supported hash type by its
// filesystem. If the filesystem's hash computation is slow or if there is any error during hash computation,
// the function returns an empty string or an error respectively.
//
// Parameters:
//   - ctx: A context to allow for timeout or cancellation of operations.
//   - object: An fs.ObjectInfo representing the object for which the hash needs to be computed.
//
// Returns:
//   - The computed hash string of the object. If the filesystem's hash computation is considered slow or
//     if there is an error in hash computation, an empty string is returned.
//   - An error if the function encounters any issues during its operation, such as failure in hash computation.
//     If no error occurs, the error is nil.
func GetHash(ctx context.Context, object fs.ObjectInfo) (string, error) {
	if object.Fs().Features().SlowHash {
		return "", nil
	}
	supportedHash := object.Fs().Hashes().GetOne()
	var hashValue string
	var err error
	if supportedHash != hash.None {
		hashValue, err = object.Hash(ctx, supportedHash)
		if err != nil {
			logger.Errorw("failed to hash", "error", err)
		}
	}

	return hashValue, err
}

var ErrStorageNotAvailable = errors.New("storage not available")

var freeSpaceWarningThreshold = 0.05
var freeSpaceErrorThreshold = 0.01

// GetRandomOutputWriter selects a storage from the provided storages list based on its available
// space and returns an associated Writer to interact with that storage.
//
// The function works as follows:
//  1. Iterates over each storage to obtain its usage information (free and total space).
//  2. Calculates a weight for each storage based on its free space ratio.
//  3. If the available space of a storage is below a threshold, it will not be considered for selection.
//  4. Selects a storage based on its weight and returns the Writer for it.
//
// Parameters:
//   - ctx: A context that allows for timeout or cancellation of operations.
//   - storages: A slice of storage models to choose from.
//
// Returns:
//   - A pointer to the ID of the selected storage.
//   - A Writer object that can be used to write data to the selected storage.
//   - An error if the function encounters any issues during its operation or if no suitable storage
//     is found. If all storages are full, it returns the ErrStorageNotAvailable error.
func GetRandomOutputWriter(ctx context.Context, storages []model.Storage) (*model.StorageID, Writer, error) {
	if len(storages) == 0 {
		return nil, nil, nil
	}

	var handlersWithWeight []struct {
		id      model.StorageID
		handler Writer
		weight  float64
	}

	for _, storage := range storages {
		handler, err := NewRCloneHandler(ctx, storage)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "failed to get storage handler for %s", storage.Name)
		}
		usage, err := handler.About(ctx)

		if usage != nil && usage.Free != nil && usage.Total != nil {
			weight := float64(*usage.Free) / float64(*usage.Total)
			if weight <= freeSpaceWarningThreshold {
				logger.Warnf("storage %s is almost full - free: %.2f%%", storage.Name, weight*100)
			}
			if weight <= freeSpaceErrorThreshold {
				logger.Errorf("storage %s is full - free: %.2f%%", storage.Name, weight*100)
				continue
			}
			handlersWithWeight = append(handlersWithWeight, struct {
				id      model.StorageID
				handler Writer
				weight  float64
			}{
				id:      storage.ID,
				handler: handler,
				weight:  float64(*usage.Free) / float64(*usage.Total),
			})
			continue
		}

		// If the getting usage request failed, it could be because the storage is not available.
		if err != nil && !errors.Is(err, ErrGetUsageNotSupported) {
			logger.Errorf("failed to get usage for storage %s: %v", storage.Name, err)
			continue
		}

		handlersWithWeight = append(handlersWithWeight, struct {
			id      model.StorageID
			handler Writer
			weight  float64
		}{
			id:      storage.ID,
			handler: handler,
			weight:  1.0,
		})
	}

	// If there is no space left in any of the output storages, return an error.
	if len(handlersWithWeight) == 0 {
		return nil, nil, ErrStorageNotAvailable
	}

	totalWeight := 0.0
	for _, item := range handlersWithWeight {
		totalWeight += item.weight
	}

	//nolint:gosec
	r := rand.Float64() * totalWeight
	for _, item := range handlersWithWeight {
		r -= item.weight
		if r <= 0 {
			return ptr.Of(item.id), item.handler, nil
		}
	}

	return nil, nil, errors.New("this line should never be reached")
}

// Open offers a file handler like interface (io.ReadSeekCloser) to an RClone path
func Open(h Handler, ctx context.Context, path string) (io.ReadSeekCloser, fs.DirEntry, error) {
	obj, err := h.Check(ctx, path)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to open object %s", path)
	}

	seeker := &rcloneSeeker{
		path:    path,
		size:    obj.Size(),
		handler: h,
	}
	return seeker, obj, nil
}

type rcloneSeeker struct {
	path    string
	size    int64
	handler Handler
	offset  int64
	file    io.ReadCloser
}

func (r *rcloneSeeker) Read(p []byte) (int, error) {
	if r.offset >= r.size {
		return 0, io.EOF
	}

	if r.file == nil {
		var err error
		r.file, _, err = r.handler.Read(context.Background(), r.path, r.offset, r.size-r.offset)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	n, err := r.file.Read(p)
	r.offset += int64(n)
	return n, err
}

func (r *rcloneSeeker) WriteTo(w io.Writer) (int64, error) {
	if r.offset >= r.size {
		return 0, io.EOF
	}

	if r.file == nil {
		var err error
		r.file, _, err = r.handler.Read(context.Background(), r.path, r.offset, r.size-r.offset)
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}

	n, err := io.Copy(w, r.file)
	r.offset += n
	return n, err
}

func (r *rcloneSeeker) Seek(offset int64, whence int) (int64, error) {
	if r.file != nil {
		err := r.file.Close()
		// Re-open file on next read.
		r.file = nil
		if err != nil {
			return 0, errors.WithStack(err)
		}
	}
	switch whence {
	case io.SeekStart:
	case io.SeekCurrent:
		offset += r.offset
	case io.SeekEnd:
		offset += r.size
	default:
		return 0, errors.New("unknown seek mode")
	}
	if offset > r.size {
		return 0, errors.New("seeking past end of file")
	}
	if offset < 0 {
		return 0, errors.New("seeking before start of file")
	}
	r.offset = offset

	return r.offset, nil
}

func (r *rcloneSeeker) Close() error {
	if r.file != nil {
		return r.file.Close()
	}
	return nil
}
