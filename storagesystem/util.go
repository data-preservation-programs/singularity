package storagesystem

import (
	"context"
	"fmt"

	"github.com/data-preservation-programs/singularity/model"
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
func IsSameEntry(ctx context.Context, file model.File, object fs.Object) (bool, string) {
	if file.Size != object.Size() {
		return false, fmt.Sprintf("size mismatch: %d != %d", file.Size, object.Size())
	}
	var err error
	// last modified can be time.Now() if fetch failed so it may not be reliable.
	// This usually won't happen for most cloud provider i.e. S3
	// Because during scanning, the modified time is already fetched.
	lastModified := object.ModTime(ctx)
	supportedHash := object.Fs().Hashes().GetOne()
	// For local file system, rclone is actually hashing the file stream which is not efficient.
	// So we skip hashing for local file system.
	// For some of the remote storage, there may not have any supported hash type.
	var hashValue string
	if supportedHash != hash.None && object.Fs().Name() != "local" {
		hashValue, err = object.Hash(ctx, supportedHash)
		if err != nil {
			logger.Errorw("failed to hash", "error", err)
		}
	}
	if file.Hash != "" && hashValue != "" && file.Hash != hashValue {
		return false, fmt.Sprintf("hash mismatch: %s != %s", file.Hash, hashValue)
	}
	return lastModified.UnixNano() == file.LastModifiedNano,
		fmt.Sprintf("last modified mismatch: %d != %d",
			lastModified.UnixNano(),
			file.LastModifiedNano)
}
