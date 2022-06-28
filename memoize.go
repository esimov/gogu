package gogu

import (
	"fmt"
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type Item struct {
	Object     any
	Expiration int64
}

type cache struct {
	duration time.Duration
	items    map[string]*Item
	mu       *sync.RWMutex
}

type Cache struct {
	*cache
}

func newCache(d time.Duration, item map[string]*Item) *cache {
	c := &cache{
		mu:       &sync.RWMutex{},
		duration: d,
		items:    item,
	}
	return c
}

func NewCache(d time.Duration) *Cache {
	items := make(map[string]*Item)
	c := newCache(d, items)

	return &Cache{c}
}

func (c *cache) Set(k string, val any, d time.Duration) error {
	c.mu.Lock()

	_, found := c.Get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item with key %s already exists. Use Update method", k)
	} else {
		c.Add(k, val, d)
	}
	c.mu.Unlock()

	return nil

}

func (c *cache) Add(k string, val any, d time.Duration) {
	var ex int64
	if d > 0 {
		ex = time.Now().Add(d).UnixNano()
	}

	c.mu.Lock()

	c.items[k] = &Item{
		Object:     val,
		Expiration: ex,
	}

	c.mu.Unlock()
}

func (c *cache) Get(k string) (*Item, bool) {
	if _, ok := c.items[k]; ok {
		return c.items[k], true
	}
	return &Item{}, false
}

func (c *cache) Update(k string, val any, d time.Duration) error {
	var ex int64
	if d > 0 {
		ex = time.Now().Add(d).UnixNano()
	}

	_, found := c.Get(k)
	if found {
		c.mu.Lock()
		c.items[k] = &Item{
			Object:     val,
			Expiration: ex,
		}
		c.mu.Unlock()

		return nil
	}

	return fmt.Errorf("Item with key %s does not exists", k)
}

func (c *cache) SetDefault(k string, val any) {
	c.Set(k, val, DefaultExpiration)
}

func (c *cache) Delete(k string) error {
	_, found := c.Get(k)
	if found {
		c.mu.Lock()
		delete(c.items, k)
		c.mu.Unlock()
	}

	return fmt.Errorf("Item with key %s does not exists", k)
}
