package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: map[string]cacheEntry{},
		mu:    sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
