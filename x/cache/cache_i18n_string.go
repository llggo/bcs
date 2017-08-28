package cache

import (
	"sync"
)

type CacheI18nEntry struct {
	ctime int64
	value map[string]string
}

const defaultLang = "en"

func (c CacheI18nEntry) Value(lang string) string {
	var v, ok = c.value[lang]
	if ok {
		return v
	}
	return c.value[defaultLang]
}

type CacheI18nFetcher func(string) (map[string]string, error)

type CacheI18nString struct {
	sync.RWMutex
	data    map[string]*CacheI18nEntry
	TTL     int32
	Fetcher CacheI18nFetcher
}

func NewCacheI18nString(fetcher CacheI18nFetcher) *CacheI18nString {
	return &CacheI18nString{
		data:    make(map[string]*CacheI18nEntry),
		TTL:     15 * 60,
		Fetcher: fetcher,
	}
}

func (c *CacheI18nString) Get(id string, lang string) (string, error) {
	var s, ok = c.data[id]
	if ok {
		return s.Value(lang), nil
	}
	var v, e = c.Fetcher(id)
	if e != nil {
		return "", e
	}
	c.Lock()
	entry := &CacheI18nEntry{value: v}
	c.data[id] = entry
	c.Unlock()
	return entry.Value(lang), nil
}
