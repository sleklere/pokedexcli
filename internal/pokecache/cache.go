package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time;
	val []byte;
}

type Cache struct {
	entries map[string]cacheEntry;
	interval time.Duration;
	mu *sync.RWMutex;
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: make(map[string]cacheEntry),
		interval: interval,
		mu: &sync.RWMutex{},
	}

	go newCache.reapLoop()

	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (cache *Cache) Remove(key string) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	delete(cache.entries, key)
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	entry, ok := cache.entries[key]
	return entry.val, ok
}

func (cache *Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for time := range ticker.C {
		// fmt.Printf("Current time: %v\n", time)

		keysToRemove := []string{}

		cache.mu.Lock()
		for key, entry := range cache.entries {
			diff := time.Sub(entry.createdAt)

			if diff > cache.interval {
				keysToRemove = append(keysToRemove, key)
			}
		}

		for _, key := range keysToRemove {
			delete(cache.entries, key)
		}

		cache.mu.Unlock()
	}
}
