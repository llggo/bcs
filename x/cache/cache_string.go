package cache

import (
	"sync"
)

type CacheEntry struct {
	ctime int64
	value string
}

type CacheFetcher func(string) (string, error)

type CacheString struct {
	sync.RWMutex
	data    map[string]*CacheEntry
	TTL     int32
	Fetcher CacheFetcher
}

func NewCacheString(fetcher CacheFetcher) *CacheString {
	return &CacheString{
		data:    make(map[string]*CacheEntry),
		TTL:     15 * 60,
		Fetcher: fetcher,
	}
}

func (c *CacheString) Get(id string) (string, error) {
	var s, ok = c.data[id]
	if ok {
		return s.value, nil
	}
	var v, e = c.Fetcher(id)
	if e != nil {
		return "", e
	}
	c.Lock()
	c.data[id] = &CacheEntry{value: v}
	c.Unlock()
	return v, nil
}
