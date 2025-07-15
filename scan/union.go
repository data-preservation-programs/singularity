package scan

import (
	"context"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/data-preservation-programs/singularity/storagesystem"
	"github.com/rs/zerolog/log"
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

	// For each upstream, compute its full path
	for _, upstream := range upstreams {
		// Handle both absolute and relative paths
		if strings.HasPrefix(upstream, "/") {
			paths[upstream] = upstream
		} else {
			paths[upstream] = filepath.Join(basePath, upstream)
		}
		log.Debug().
			Str("upstream", upstream).
			Str("fullPath", paths[upstream]).
			Msg("Mapped upstream to full path")
	}

	return paths
}
