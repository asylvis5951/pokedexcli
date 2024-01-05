package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entry    map[string]entry
	interval time.Duration
}

type entry struct {
	createdAt time.Time
	value     []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = entry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.entry[key]; !ok {
		return []byte{}, false
	}
	return c.entry[key].value, true
}

func (c *Cache) Delete(k string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entry, k)
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case t := <-ticker.C:
			// if interval is greater than age of entry, delete entry.
			for k, v := range c.entry {
				if t.Sub(v.createdAt) > c.interval {
					c.Delete(k)
				}
			}
		}
	}
}

func NewCache(i time.Duration) *Cache {
	ent := map[string]entry{}
	c := Cache{
		entry:    ent,
		interval: i,
	}
	go c.reapLoop()
	return &c
}
