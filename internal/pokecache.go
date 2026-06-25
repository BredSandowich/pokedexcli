package pokecache

import (
	"time"
	"sync"
)


type Cache struct {
	mu sync.Mutex
	mem map[string]cacheEntry
}


type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	//
	c := Cache{}
	c.mem = make(map[string]cacheEntry)
	
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ce := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	c.mem[key] = ce
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.mem[key]; ok {
		return val.val, true
	}
	return nil, false
}