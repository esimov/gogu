package main

import (
	"fmt"
	"time"

	"github.com/esimov/torx"
	"github.com/esimov/torx/cache"
)

func main() {
	sampleMap := map[string]any{
		"foo": "one",
		"bar": "two",
		"baz": "three",
	}

	m := torx.NewMemoizer[string, any](time.Second, time.Minute)

	expensiveOp := func() (*cache.Item[any], error) {
		// Here we are simulating an expensive operation.
		time.Sleep(500 * time.Millisecond)

		foo := torx.FindByKey(sampleMap, func(key string) bool {
			return key == "foo"
		})
		m.Cache.MapToCache(foo, cache.DefaultExpiration)

		item, err := m.Cache.Get("foo")
		if err != nil {
			return nil, err
		}
		return item, nil
	}

	// At the first call key1 does not exists, so the simulated extensive operation will take more time.
	data, _ := m.Memoize("item", expensiveOp)
	fmt.Println(data)

	// Second time the date will be read from cache. It will return the result instantly.
	data, _ = m.Memoize("item", expensiveOp)
	fmt.Println(data)
}
