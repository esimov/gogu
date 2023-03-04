// Package cache implements a basic in memory key-value storage system using map as storing mechanism.
// The cache and the cache items also have an expiration time. The cache will be invalidated
// once the expiration time is reached. On cache initialization a cleanup interval is also required.
// The scope of the cleanup method is to run at a predefined interval and to remove all the expired cache items.
package cache

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

// Item holds the cache object (which could be of any type) and an expiration time.
// The expiration time defines the object lifetime.
type Item[V any] struct {
	object     V
	expiration int64
}

type cache[K ~string, V any] struct {
	mu         sync.RWMutex
	items      map[K]*Item[V]
	done       chan struct{}
	expTime    time.Duration
	cleanupInt time.Duration
}

// Cache is a publicly available struct type, which incorporates the
// unexported cache struct type holding the cache components.
type Cache[K ~string, V any] struct {
	*cache[K, V]
}

// newCache has a local scope only. `New` will be used for the cache instantiation outside this package.
func newCache[K ~string, V any](expTime, cleanupInt time.Duration, item map[K]*Item[V]) *cache[K, V] {
	c := &cache[K, V]{
		mu:         sync.RWMutex{},
		items:      item,
		expTime:    expTime,
		cleanupInt: cleanupInt,
		done:       make(chan struct{}),
	}
	return c
}

// New instantiates a cache struct which requires an expiration time and a cleanup interval.
// The cache will be invalidated once the expiration time is reached.
// If the expiration time is less than zero (or NoExpiration) the cache items will never expire and should be deleted manually.
// A cleanup method is running in the background and removes the expired caches at a predefined interval.
func New[K ~string, V any](expTime, cleanupTime time.Duration) *Cache[K, V] {
	items := make(map[K]*Item[V])
	c := newCache(expTime, cleanupTime, items)

	if cleanupTime > 0 {
		go c.cleanup()
		// We need to make sure that the goroutine responsible for the cache eviction stops after the cleanup.
		// This is the reason why runtime.SetFinalizer is used.
		// This method is invoked when the garbage collector finds an unreachable block ready to be collected.
		runtime.SetFinalizer(c, stopCleanup[K, V])
	}

	return &Cache[K, V]{c}
}

// Set inserts a new item into the cache, but first verifies if an item with the same key already exists in the cache.
// In case an item with the specified key already exists in the cache it will return an error.
func (c *Cache[K, V]) Set(key K, val V, d time.Duration) error {
	item, err := c.Get(key)
	if item != nil && err == nil {
		return fmt.Errorf("item with key '%v' already exists. Use the Update method", key)
	}
	c.add(key, val, d)

	return nil
}

// SetDefault adds a new item into the cache with the default expiration time.
func (c *Cache[K, V]) SetDefault(key K, val V) error {
	return c.Set(key, val, DefaultExpiration)
}

// add inserts a new item into the cache together with an expiration time.
// If the duration is 0 (or DefaultExpiration) the cache default expiration time is used.
// If the duration is < 0 (or NoExpiration), the item never expires and should be removed manually.
func (c *Cache[K, V]) add(key K, val V, d time.Duration) error {
	var exp int64

	if d == DefaultExpiration {
		d = c.expTime
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
	} else if d < 0 {
		exp = int64(NoExpiration)
	}

	item, err := c.Get(key)
	if item != nil && err != nil {
		return fmt.Errorf("item with key '%v' already exists", key)
	}

	switch any(val).(type) {
	case string:
		if len(any(val).(string)) == 0 {
			return fmt.Errorf("value of type string cannot be empty")
		}
	}

	c.mu.Lock()
	c.items[key] = &Item[V]{
		object:     val,
		expiration: exp,
	}
	c.mu.Unlock()

	return nil
}

// Get returns a cache item defined by its key. If the item is expired an error is returned.
// If an item is expired it's considered as nonexistent, it will be evicted from the cache
// when the purge method is invoked at the predefined interval.
func (c *Cache[K, V]) Get(key K) (*Item[V], error) {
	c.mu.RLock()
	if item, ok := c.items[key]; ok {
		if item.expiration > 0 {
			now := time.Now().UnixNano()
			if now > item.expiration {
				c.mu.RUnlock()
				return nil, fmt.Errorf("item with key '%v' expired", key)
			}
		}
		c.mu.RUnlock()
		return item, nil
	}
	c.mu.RUnlock()
	return nil, fmt.Errorf("item with key '%v' not found", key)
}

// Val returns the effective value of the cache item.
func (it *Item[V]) Val() V {
	var v V
	if it != nil {
		return it.object
	}
	return v
}

// Update replaces a cache item with the new value.
func (c *Cache[K, V]) Update(key K, val V, d time.Duration) error {
	item, err := c.Get(key)
	if item != nil && err != nil {
		return err
	}
	return c.add(key, val, d)
}

// Delete removes a cache item.
func (c *Cache[K, V]) Delete(key K) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.delete(key)
}

// delete has a local scope only.
func (c *cache[K, V]) delete(key K) error {
	if _, ok := c.items[key]; ok {
		delete(c.items, key)

		return nil
	}

	return fmt.Errorf("item with key '%v' does not exists", key)
}

// DeleteExpired removes all the expired items from the cache.
func (c *cache[K, V]) DeleteExpired() error {
	var err error

	now := time.Now().UnixNano()

	c.mu.Lock()
	for k, item := range c.items {
		if now > item.expiration && item.expiration != int64(NoExpiration) {
			if e := c.delete(k); e != nil {
				err = errors.Join(err, e)
			}
		}

	}
	c.mu.Unlock()

	return errors.Unwrap(err)
}

// Flush removes all the existing items in the cache.
func (c *Cache[K, V]) Flush() {
	c.mu.Lock()
	c.items = make(map[K]*Item[V])
	c.mu.Unlock()
}

// List returns the cache items which are not expired.
func (c *Cache[K, V]) List() map[K]*Item[V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.items
}

// Count returns the number of existing items in the cache.
func (c *Cache[K, V]) Count() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()

	return n
}

// MapToCache transfers the map values into the cache.
func (c *Cache[K, V]) MapToCache(m map[K]V, d time.Duration) error {
	var err error

	for k, v := range m {
		e := c.Set(k, v, d)
		err = errors.Join(err, e)
	}

	return errors.Unwrap(err)
}

// IsExpired checks if a cache item is expired.
func (c *Cache[K, V]) IsExpired(key K) bool {
	item, err := c.Get(key)
	if item != nil && err != nil {
		if item.expiration > time.Now().UnixNano() {
			return true
		}
	}
	return false
}

// cleanup runs the cache cleanup function at the specified time interval an removes all the expired cache items.
func (c *cache[K, V]) cleanup() {
	tick := time.NewTicker(c.cleanupInt)

	for {
		select {
		case <-tick.C:
			c.DeleteExpired()
		case <-c.done:
			tick.Stop()
			return
		}
	}
}

// stopCleanup stops the cleanup process once the cache item goes out of scope and became unreachable.
func stopCleanup[K ~string, V any](c *cache[K, V]) {
	c.done <- struct{}{}
}
