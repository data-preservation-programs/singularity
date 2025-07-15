package scan

import (
	"context"
	"path/filepath"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/storagesystem"
)

// GetUnionUpstreams enumerates the upstream directories in a union storage configuration
func GetUnionUpstreams(ctx context.Context, storage *storagesystem.RCloneHandler) ([]string, error) {
	var upstreams []string
	
	// Check if the storage is a union backend
	if fs, ok := storage.Fs().(interface{ ListUpstreams() []string }); ok {
		upstreams = fs.ListUpstreams()
	} else {
		return nil, errors.New("storage is not a union backend or does not support listing upstreams")
	}

	return upstreams, nil
}

// GetUpstreamPaths returns the full paths for each upstream in a union storage
func GetUpstreamPaths(storage *storagesystem.RCloneHandler, upstreams []string) map[string]string {
	paths := make(map[string]string)
	basePath := storage.Fs().Root()

	for _, upstream := range upstreams {
		paths[upstream] = filepath.Join(basePath, upstream)
	}

	return paths
}
