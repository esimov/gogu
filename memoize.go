package gogu

import (
	"time"

	"github.com/esimov/gogu/cache"
	"golang.org/x/sync/singleflight"
)

// Memoizer is a two component struct type used to memoize the results of a function execution.
// It holds an exported Cache storage and a singleflight group which is used
// to guarantee that only one function execution is in flight for a given key.
type Memoizer[T ~string, V any] struct {
	Cache *cache.Cache[T, V]
	group *singleflight.Group
}

// NewMemoizer instantiates a new Memoizer.
func NewMemoizer[T ~string, V any](expiration, cleanup time.Duration) *Memoizer[T, V] {
	return &Memoizer[T, V]{
		Cache: cache.New[T, V](expiration, cleanup),
		group: &singleflight.Group{},
	}
}

// Memoize returns the item under a specific key instantly in case the key exists,
// otherwise returns the results of the given function, making sure that only one execution
// is in-flight for a given key at a time.
//
// This method is useful for caching the result of a time-consuming operation when is more important
// to return a slightly outdated result, than to wait for an operation to complete before serving it.
func (m Memoizer[T, V]) Memoize(key T, fn func() (*cache.Item[V], error)) (*cache.Item[V], error) {
	item, _ := m.Cache.Get(key)
	if item != nil {
		return item, nil
	}

	data, err, _ := m.group.Do(string(key), func() (any, error) {
		item, err := fn()
		if err == nil {
			m.Cache.SetDefault(key, item.Val())
		}
		return item, err
	})

	return data.(*cache.Item[V]), err
}
