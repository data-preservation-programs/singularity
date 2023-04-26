package pack

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/disk"
)

func GetPathWithMostSpace(paths []string) (string, error) {
	var maxSpace uint64
	var maxPath string

	for _, path := range paths {
		usage, err := disk.Usage(path)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get disk usage for path %s", path)
		}

		availableSpace := usage.Free

		if availableSpace > maxSpace {
			maxSpace = availableSpace
			maxPath = path
		}
	}

	if maxPath == "" {
		return "", errors.New("no paths provided")
	}

	return maxPath, nil
}
