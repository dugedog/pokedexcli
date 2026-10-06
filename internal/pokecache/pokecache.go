package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheMap map[string]cacheEntry
	duration int
	mu       sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	intervalSeconds := interval * time.Second
	createdCache := &Cache{duration: int(intervalSeconds)}
	tracker := time.NewTicker(intervalSeconds)
	go func() {
		for loop := range tracker.C {
			print(loop)
			createdCache.reapLoop(intervalSeconds)
		}
	}()
	return createdCache

}
func (c *Cache) Add(key string, newVal []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{createdAt: time.Now(),
		val: newVal}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	lookUpValue, ok := c.cacheMap[key]
	return lookUpValue.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.cacheMap {
		if value.createdAt.Add(interval).Before(time.Now()) {
			delete(c.cacheMap, key)
		}

	}
}
