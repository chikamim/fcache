package fcache

import "io"

type CacheValue struct {
	Value []byte
}

func (mv CacheValue) Len() int64 {
	return int64(len(mv.Value))
}

// Cache cache interface definition.The cache can be memory cache,
// disk cache or net cache. We implementation cache with LRU algorithm.
// Note: disk cache NOT support 'extra' field in current version.
type Cache interface {
	// Set set cache with key-value pair.
	Set(key string, value []byte, extra ...interface{}) error

	// Get get cache with key, nil will be return if key is not exist.
	Get(key string) (value []byte, extra interface{}, err error)

	// Reader returns cache reader, nil will be return if key is not exist.
	Reader(key string) (r io.ReadCloser, extra interface{}, err error)

	// GetHitInfo get cache hit info, return the count of hit visitor
	// and the count of total visitor
	GetHitInfo() (hitCount, totalCount int64)

	// Clear clear cache with key
	Clear(key string) error

	// ClearAll clear all cache
	ClearAll() error
}

// NewMemCache return a memory cache instance.
// maxSize: memory cache max size.
// needCryptKey:if set true, crypt key(with md5), if key is too long,
// crypt key maybe better choice.
// ttl:time to live of cache data(second).
func NewMemCache(maxSize int64, needCryptKey bool, ttl ...int64) Cache {
	if len(ttl) > 0 {
		return newMemCache(maxSize, needCryptKey, ttl[0])
	}
	return newMemCache(maxSize, needCryptKey, 0)
}

// NewDiskCache return a disk cache instance.Params of maxSize:disk cache
// max size.Param of needCryptKey:if set true, crypt key(with md5), if key
// is too long, crypt key maybe better choice.Param of cacheDir:disk cache
// directory.Param of ttl:time to live of cache data(second).
func NewDiskCache(maxSize int64, needCryptKey bool, cacheDir string, ttl ...int64) Cache {
	if len(ttl) > 0 {
		return newDiskCache(maxSize, needCryptKey, cacheDir, ttl[0])
	}
	return newDiskCache(maxSize, needCryptKey, cacheDir, 0)
}
