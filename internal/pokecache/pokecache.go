package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	register	map[string]cacheEntry
	mu			*sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache {
		register:	map[string]cacheEntry{},
		mu:			&sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val:	val,
	}
	c.mu.Lock()
	c.register[key] = entry
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, exists := c.register[key]
	c.mu.Unlock()
	if !exists {
		return entry.val, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		var entriesToDelete []string
		t := <- ticker.C
		c.mu.Lock()
		for key, entry := range c.register {
				if t.Sub(entry.createdAt) > interval {
					entriesToDelete = append(entriesToDelete, key)
				}
			}
		for _, entry := range entriesToDelete {
			delete(c.register, entry)
		}
		c.mu.Unlock()
	}
}
