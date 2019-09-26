package cache

import (
	"errors"
	"strconv"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prysmaticlabs/prysm/shared/params"
	"k8s.io/client-go/tools/cache"
)

var (
	// ErrNotCommitteeItem will be returned when a cache object is not a pointer to
	// a CommitteeByShardEpoch struct.
	ErrNotCommitteeItem = errors.New("object is not a committee info")

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
	StartShard     uint64
	CommitteeCount uint64
	Epoch          uint64
	Committee      []uint64
}

// CommitteeCache is a struct with 1 queue for looking up committee list by epoch and shard.
type CommitteeCache struct {
	CommitteeCache *cache.FIFO
	lock           sync.RWMutex
}

// committeeKeyFn takes the epoch as the key for the active indices of a given epoch.
func committeeKeyFn(obj interface{}) (string, error) {
	info, ok := obj.(*CommitteeByShardEpoch)
	if !ok {
		return "", ErrNotCommitteeItem
	}

	return strconv.Itoa(int(info.Epoch)), nil
}

// NewCommitteeCache creates a new committee cache for storing/accessing committees.
func NewCommitteeCache() *CommitteeCache {
	return &CommitteeCache{
		CommitteeCache: cache.NewFIFO(committeeKeyFn),
	}
}

// CommitteeInEpochShard fetches CommitteeByShardEpoch by epoch and shard. Returns true with a
// reference to the CommitteeByShardEpoch info, if exists. Otherwise returns false, nil.
func (c *CommitteeCache) CommitteeInEpochShard(epoch uint64, shard uint64) ([]uint64, error) {
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

	item, ok := obj.(*CommitteeByShardEpoch)
	if !ok {
		return nil, ErrNotCommitteeItem
	}

	start, end := startEndIndices(item, shard)

	return item.Committee[start:end], nil
}

// AddCommitteeList adds CommitteeByShardEpoch object to the cache. This method also trims the least
// recently added CommitteeByShardEpoch object if the cache size has ready the max cache size limit.
func (c *CommitteeCache) AddCommitteeList(committee *CommitteeByShardEpoch) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if err := c.CommitteeCache.AddIfNotPresent(committee); err != nil {
		return err
	}
	trim(c.CommitteeCache, maxCommitteeListSize)
	return nil
}

// Epochs returns the stored epochs in the committee cache.
func (c *CommitteeCache) Epochs() ([]uint64, error) {
	epochs := make([]uint64, len(c.CommitteeCache.ListKeys()))
	for i, s := range c.CommitteeCache.ListKeys() {
		epoch, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		epochs[i] = uint64(epoch)
	}
	return epochs, nil
}

// EpochInCache returns true if the input epoch is in cache.
func (c *CommitteeCache) EpochInCache(wantedEpoch uint64) (bool, error) {
	for _, s := range c.CommitteeCache.ListKeys() {
		epoch, err := strconv.Atoi(s)
		if err != nil {
			return false, err
		}
		if wantedEpoch == uint64(epoch) {
			return true, nil
		}
	}
	return false, nil
}

func startEndIndices(c *CommitteeByShardEpoch, wantedShard uint64) (uint64, uint64) {
	shardCount := params.BeaconConfig().ShardCount
	currentShard := (wantedShard + shardCount - c.StartShard) % shardCount
	validatorCount := uint64(len(c.Committee))
	start := (validatorCount * currentShard) / c.CommitteeCount
	end := validatorCount * (currentShard + 1) / c.CommitteeCount

	return start, end
}
