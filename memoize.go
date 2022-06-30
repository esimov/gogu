package gogu

import (
	"fmt"
	"time"

	"golang.org/x/sync/singleflight"
)

type Memoizer[T ~string, V any] struct {
	Cache *Cache[T, V]
	group *singleflight.Group
}

func NewMemoizer[T ~string, V any](expiration, cleanup time.Duration) *Memoizer[T, V] {
	return &Memoizer[T, V]{
		Cache: NewCache[T, V](expiration, cleanup),
		group: &singleflight.Group{},
	}
}

func (m Memoizer[T, V]) Memoize(key T, fn func() (*Item[V], error)) (*Item[V], error) {
	item, _ := m.Cache.Get(key)
	if item != nil {
		return item, nil
	}

	data, err, _ := m.group.Do(string(key), func() (any, error) {
		item, err := fn()
		if err == nil {
			m.Cache.Set(key, item.Object, DefaultExpiration)
		}
		return item, err
	})
	fmt.Println("Data:", data)

	return data.(*Item[V]), err
}
