package storage

import (
	"log"
	"strings"
	"sync"

	"github.com/Kanbenn/mywbgonats/internal/models"
)

type Cache struct {
	m  map[string][]byte
	mu *sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		make(map[string][]byte),
		new(sync.RWMutex)}
}

func (c *Cache) Add(key string, val []byte) {
	if len(key) < 1 {
		log.Println("Cache.Add error: empty key", key)
		return
	}
	c.withFullLock(func() {
		_, exists := c.m[key]
		if exists {
			log.Println("Cache.Add error: key already exists", key)
			return
		}
		c.m[key] = val
	})
}

func (c *Cache) AddBatch(in []models.Order) {
	c.withFullLock(func() {
		for _, o := range in {
			c.m[o.ID] = o.Data
		}
	})
}

func (c *Cache) Get(key string) (data []byte, found bool) {
	c.withRLock(func() {
		data, found = c.m[key]
	})
	return data, found
}

func (c *Cache) GetAllKeys() (keys []string) {
	c.withRLock(func() {
		for key := range c.m {
			keys = append(keys, key)
		}
	})
	return keys
}

func (c *Cache) withRLock(f func()) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	f()
}

func (c *Cache) withFullLock(f func()) {
	c.mu.Lock()
	defer c.mu.Unlock()
	f()
}

func (c *Cache) String() string {
	var sb strings.Builder
	c.withRLock(func() {
		for k := range c.m {
			sb.WriteString("key: " + k)
		}
	})
	return sb.String()
}
