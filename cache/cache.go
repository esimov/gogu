package cache

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"go.uber.org/multierr"
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

type cache[T ~string, V any] struct {
	mu         *sync.RWMutex
	items      map[T]*Item[V]
	expTime    time.Duration
	cleanupInt time.Duration
	done       chan struct{}
}

// Exported cache struct.
type Cache[T ~string, V any] struct {
	*cache[T, V]
}

// newCache has a local scope only. NewCache will be used for the cache instantiation outside of this package.
func newCache[T ~string, V any](expTime, cleanupInt time.Duration, item map[T]*Item[V]) *cache[T, V] {
	c := &cache[T, V]{
		mu:         &sync.RWMutex{},
		items:      item,
		expTime:    expTime,
		cleanupInt: cleanupInt,
		done:       make(chan struct{}),
	}
	return c
}

// New instantiates a cache struct which requires an expiration time and a cleanup interval.
// The cache will be invalidated once the expiration time is reached.
// If the expiration time is less than zero (or NoExpiration) the cache items will never expire and should be manually deleted.
// A cleanup method is running in the background and removes the expired caches at a predifined interval.
func New[T ~string, V any](expTime, cleanupTime time.Duration) *Cache[T, V] {
	items := make(map[T]*Item[V])
	c := newCache(expTime, cleanupTime, items)

	if cleanupTime > 0 {
		go c.cleanup()
		// In case no human interaction is happening in the background, we need to make sure
		// that the goroutine responsible for the cache purge stops after the cleanup.
		// This is the reason why runtime.SetFinalizer is used.
		// This method is invoked when the garbage collector finds an unreachable block ready to be collected.
		runtime.SetFinalizer(c, stopCleanup[T, V])
	}

	return &Cache[T, V]{c}
}

// Set add a new item to the cache, but first verifies if an item with the same key already exists in the cache.
func (c *cache[T, V]) Set(key T, val V, d time.Duration) error {
	item, err := c.Get(key)
	if item != nil && err == nil {
		return fmt.Errorf("item with key '%v' already exists. Use the Update method", key)
	}
	c.add(key, val, d)

	return nil
}

// SetDefault includes a new item in the cache with the default expiration time.
func (c *cache[T, V]) SetDefault(key T, val V) error {
	return c.Set(key, val, DefaultExpiration)
}

// add place a new item into the cache. This method is not exported.
// It puts the item into the cache together with the expiration time.
// If the duration is 0 (or DefaultExpiration) the cache default expiration time is used.
// If the duration is < 0 (or NoExpiration), the item never expires and should be manually deleted.
func (c *cache[T, V]) add(key T, val V, d time.Duration) error {
	var exp int64

	if d == DefaultExpiration {
		d = c.expTime
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
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

// Get returns the item from the cache identifiable by its key.
// If an item expiration time is reached an error is returned instead of the item itself.
// Every time the item will be purged by the cleanup method at the predifined interval.
func (c *cache[T, V]) Get(key T) (*Item[V], error) {
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

// Update replaces an item from the cache with the new values.
func (c *cache[T, V]) Update(key T, val V, d time.Duration) error {
	item, err := c.Get(key)
	if item != nil && err != nil {
		return err
	}
	return c.add(key, val, d)
}

// Delete removes an item from the cache.
func (c *cache[T, V]) Delete(key T) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.delete(key)
}

// delete has a local scope only.
func (c *cache[T, V]) delete(key T) error {
	if _, ok := c.items[key]; ok {
		delete(c.items, key)

		return nil
	}

	return fmt.Errorf("item with key '%v' does not exists", key)
}

// DeleteExpired removes all the expired items from the cache.
func (c *cache[T, V]) DeleteExpired() error {
	var err error

	now := time.Now().UnixNano()

	c.mu.Lock()
	// TODO replace multierr package when proposal https://github.com/golang/go/issues/53435
	// will be accepted and integrated into the standard library.
	for k, item := range c.items {
		if now > item.expiration {
			if e := c.delete(k); e != nil {
				err = multierr.Append(err, e)
			}
		}

	}
	c.mu.Unlock()

	return multierr.Combine(err)
}

// Flush deletes all the items from the cache.
func (c *cache[T, V]) Flush() {
	c.mu.Lock()
	c.items = make(map[T]*Item[V])
	c.mu.Unlock()
}

// List returns the cache items.
func (c *cache[T, V]) List() map[T]*Item[V] {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.items
}

// Count the number of items existing in the cache.
func (c *cache[T, V]) Count() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()

	return n
}

// MapToCache moves the items from a map into the cache.
func (c *cache[T, V]) MapToCache(m map[T]V, d time.Duration) error {
	var err error

	for k, v := range m {
		e := c.Set(k, v, d)
		err = multierr.Append(err, e)
	}

	return multierr.Combine(err)
}

// IsExpired checks if an item is expired or not.
func (c *cache[T, V]) IsExpired(key T) bool {
	item, err := c.Get(key)
	if item != nil && err != nil {
		if item.expiration > time.Now().UnixNano() {
			return true
		}
	}
	return false
}

// cleanup runs the cache cleanup function at the specified time interval an removes all the expired cache items.
func (c *cache[T, V]) cleanup() {
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
func stopCleanup[T ~string, V any](c *cache[T, V]) {
	c.done <- struct{}{}
}
