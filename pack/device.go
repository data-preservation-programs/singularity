package pack

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/jellydator/ttlcache/v3"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/disk"
	"math/rand"
	"time"
)

var deviceCache = ttlcache.New[string, *disk.UsageStat](
	ttlcache.WithTTL[string, *disk.UsageStat](time.Hour))

func getRandomStringByWeight(pathMap map[string]uint64) string {
	var totalWeight uint64
	for _, weight := range pathMap {
		totalWeight += weight
	}

	rand.Seed(time.Now().UnixNano())
	randWeight := rand.Intn(int(totalWeight))

	var currentWeight uint64
	for path, weight := range pathMap {
		currentWeight += weight
		if randWeight <= int(currentWeight) {
			return path
		}
	}

	return ""
}

func getUsage(path string) (*disk.UsageStat, error) {
	item := deviceCache.Get(path)
	if item != nil && !item.IsExpired() {
		return item.Value(), nil
	}

	logging.Logger("device").Debugf("getting disk usage for path %s", path)
	usage, err := disk.Usage(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get disk usage for path %s", path)
	}
	logging.Logger("device").With("usage", *usage).Debugf("got disk usage for path %s", path)
	deviceCache.Set(path, usage, ttlcache.DefaultTTL)
	return usage, nil
}

func GetPathWithMostSpace(paths []string) (string, error) {
	var pathMap = make(map[string]uint64)

	for _, path := range paths {
		usage, err := getUsage(path)
		if err != nil {
			return "", errors.Wrapf(err, "failed to get disk usage for path %s", path)
		}

		availableSpace := usage.Free
		pathMap[path] = availableSpace
	}

	if len(pathMap) == 0 {
		return "", errors.New("no paths provided")
	}

	// Get a random path from the list of paths with the most space
	return getRandomStringByWeight(pathMap), nil
}
