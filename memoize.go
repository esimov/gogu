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

type Item[V any] struct {
	Object     V
	Expiration int64
}

type cache[T ~string, V any] struct {
	mu         *sync.RWMutex
	items      map[T]*Item[V]
	expiration time.Duration
}

type Cache[T ~string, V any] struct {
	*cache[T, V]
}

func newCache[T ~string, V any](du time.Duration, item map[T]*Item[V]) *cache[T, V] {
	c := &cache[T, V]{
		mu:         &sync.RWMutex{},
		items:      item,
		expiration: du,
	}
	return c
}

func NewCache[T ~string, V any](d time.Duration) *Cache[T, V] {
	items := make(map[T]*Item[V])
	c := newCache(d, items)

	return &Cache[T, V]{c}
}

func (c *cache[T, V]) Set(key T, val V, d time.Duration) error {
	_, err := c.Get(key)
	if err != nil {
		return fmt.Errorf("item with key %v already exists. Use the Update method", key)
	}
	c.Add(key, val, d)

	return nil
}

func (c *cache[T, V]) Add(key T, val V, d time.Duration) error {
	var exp int64

	if d == DefaultExpiration {
		d = c.expiration
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
	}

	_, err := c.Get(key)
	if err != nil {
		return fmt.Errorf("item with key %v already exists", key)
	}

	c.mu.Lock()
	c.items[key] = &Item[V]{
		Object:     val,
		Expiration: exp,
	}
	c.mu.Unlock()

	return nil
}

func (c *cache[T, V]) Get(key T) (*Item[V], error) {
	c.mu.RLock()
	if item, ok := c.items[key]; ok {
		if item.Expiration > 0 {
			now := time.Now().UnixNano()
			if now > item.Expiration {
				c.mu.RUnlock()
				return &Item[V]{}, fmt.Errorf("item with key %v expired", key)
			}
		}
		c.mu.RUnlock()
		return c.items[key], nil
	}
	c.mu.RUnlock()
	return &Item[V]{}, nil
}

func (c *cache[T, V]) Update(key T, val V, d time.Duration) error {
	c.mu.Lock()
	_, err := c.Get(key)
	if err != nil {
		c.mu.Unlock()
		return fmt.Errorf("item with key %v does not exists", key)
	}
	c.Set(key, val, d)
	c.mu.Unlock()

	return nil
}

func (c *cache[T, V]) SetDefault(key T, val V) {
	c.Set(key, val, DefaultExpiration)
}

func (c *cache[T, V]) Delete(key T) error {
	_, err := c.Get(key)
	if err != nil {
		c.mu.Lock()
		delete(c.items, key)
		c.mu.Unlock()
	}

	return fmt.Errorf("item with key %v does not exists", key)
}

func (c *cache[T, V]) IsExpired(key T) bool {
	item, err := c.Get(key)
	if err != nil {
		if item.Expiration > time.Now().UnixNano() {
			return true
		}
	}
	return false
}
