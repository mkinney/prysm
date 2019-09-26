package cache

import (
	"errors"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"k8s.io/client-go/tools/cache"
)

var (
	// ErrNotCommitteeInfo will be returned when a cache object is not a pointer to
	// a CommitteeByShardEpoch struct.
	ErrNotCommitteeInfo = errors.New("object is not a committee info")

	// maxCommitteeInfoSize defines the max number of committee info can cache.
	// 2 for current epoch and next epoch.
	maxCommitteeListSize = 2

	// Metrics.
	CommitteeCacheMiss = promauto.NewCounter(prometheus.CounterOpts{
		Name: "committee_cache_miss",
		Help: "The number of committee requests that aren't present in the cache.",
	})
	CommitteeCacheHit = promauto.NewCounter(prometheus.CounterOpts{
		Name: "committee_cache_hit",
		Help: "The number of committee requests that are present in the cache.",
	})
)

// CommitteeByShardEpoch defines the committee per epoch.
type CommitteeByShardEpoch struct {
	Shard uint64
	Epoch         uint64
	Committee []uint64
}

// CommitteeCache is a struct with 1 queue for looking up committee list by epoch and shard.
type CommitteeCache struct {
	CommitteeCache *cache.FIFO
	lock           sync.RWMutex
}

// CommitteeKeyFn takes the epoch as the key for the active indices of a given epoch.
func CommitteeKeyFn(obj interface{}) (string, error) {
	info, ok := obj.(*CommitteeByShardEpoch)
	if !ok {
		return "", ErrNotCommitteeInfo
	}

	return strconv.Itoa(int(info.Epoch)), nil
}

// NewCommitteeCache creates a new committee cache for storing/accessing committees.
func NewCommitteeCache() *CommitteeCache {
	return &CommitteeCache{
		CommitteeCache: cache.NewFIFO(CommitteeKeyFn),
	}
}

// CommitteeInEpoch fetches CommitteeByShardEpoch by epoch. Returns true with a
// reference to the CommitteeInEpoch info, if exists. Otherwise returns false, nil.
func (c *CommitteeCache) CommitteeInEpoch(epoch uint64) ([]uint64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	obj, exists, err := c.CommitteeCache.GetByKey(strconv.Itoa(int(epoch)))
	if err != nil {
		return nil, err
	}

	if exists {
		CommitteeCacheHit.Inc()
	} else {
		CommitteeCacheMiss.Inc()
		return nil, nil
	}

	aInfo, ok := obj.(*CommitteeByShardEpoch)
	if !ok {
		return nil, ErrNotCommitteeInfo
	}

	return aInfo.Committee, nil
}

// AddCommitteeList adds CommitteeByShardEpoch object to the cache. This method also trims the least
// recently added CommitteeByShardEpoch object if the cache size has ready the max cache size limit.
func (c *CommitteeCache) AddCommitteeList(Committee *CommitteeByShardEpoch) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if err := c.CommitteeCache.AddIfNotPresent(Committee); err != nil {
		return err
	}

	trim(c.CommitteeCache, maxCommitteeListSize)
	return nil
}

// CommitteeKeys returns the keys of the active indices cache.
func (c *CommitteeCache) CommitteeKeys() []string {
	return c.CommitteeCache.ListKeys()
}
