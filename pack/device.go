package pack

import (
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/disk"
	"math/rand"
	"time"
)

func getRandomString(strings []string) string {
	rand.Seed(time.Now().UnixNano())
	// nolint:gosec
	randomIndex := rand.Intn(len(strings))
	return strings[randomIndex]
}

func GetPathWithMostSpace(paths []string) (string, error) {
	var maxSpace uint64
	var maxPaths []string

	for _, path := range paths {
		usage, err := disk.Usage(path)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get disk usage for path %s", path)
		}

		availableSpace := usage.Free

		if availableSpace == maxSpace {
			maxPaths = append(maxPaths, path)
		} else if availableSpace > maxSpace {
			maxSpace = availableSpace
			maxPaths = []string{path}
		}
	}

	if len(maxPaths) == 0 {
		return "", errors.New("no paths provided")
	}

	// Get a random path from the list of paths with the most space
	return getRandomString(maxPaths), nil
}
