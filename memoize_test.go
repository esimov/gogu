package gogu

import (
	"testing"
	"time"

	"github.com/esimov/gogu/cache"
	"github.com/stretchr/testify/assert"
)

func TestMemoize(t *testing.T) {
	assert := assert.New(t)

	m := NewMemoizer[string, any](time.Second, time.Minute)

	sampleItem := map[string]any{
		"foo": "one",
		"bar": "two",
		"baz": "three",
	}

	// Here we are simulating an expensive operation.
	expensiveOp := func() (*cache.Item[any], error) {
		// Here we are simulating an expensive operation.
		time.Sleep(500 * time.Millisecond)

		foo := FindByKey(sampleItem, func(key string) bool {
			return key == "foo"
		})
		m.Cache.MapToCache(foo, cache.DefaultExpiration)

		item, err := m.Cache.Get("foo")
		if err != nil {
			return nil, err
		}
		return item, nil
	}
	assert.Empty(m.Cache.List())
	// Caching the result of some expensive fictive operation result.
	data, err := m.Memoize("key1", expensiveOp)
	assert.NoError(err)
	assert.NotEmpty(data)
	assert.NotEmpty(m.Cache.List())
	item, _ := m.Cache.Get("key1")
	assert.Equal("one", item.Val())

	// Serving the expensive operation result from the cache. This should return instantly.
	// If it would invoked the expensiveOp function this would be introduced a 500 millisecond latency.
	start := time.Now()
	data, err = m.Memoize("key1", expensiveOp)
	after := time.Since(start).Milliseconds()

	assert.NoError(err)
	assert.NotEmpty(data)
	assert.LessOrEqual(int(after), 1)
}
