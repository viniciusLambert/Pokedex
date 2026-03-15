package pokecache

import "time"

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		c.reap(time.Now().UTC(), duration)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.entry {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.entry, key)
		}
	}
}
