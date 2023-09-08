package storagesystem

import (
	"context"
	"fmt"
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
// - In certain cases (e.g., failures during fetch), the last modified time might not be reliable.
// - For local file systems, hash computation is skipped to avoid inefficient operations.
func IsSameEntry(ctx context.Context, file model.File, object fs.ObjectInfo) (bool, string) {
	if file.Size != object.Size() {
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

func GetRandomOutputWriter(ctx context.Context, storages []model.Storage) (*uint32, Writer, error) {
	if len(storages) == 0 {
		return nil, nil, nil
	}

	var handlersWithWeight []struct {
		id      uint32
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
				id      uint32
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
			id      uint32
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
