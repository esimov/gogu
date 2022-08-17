package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	sampleMap := map[string]any{
		"foo": "one",
		"bar": "two",
		"baz": "three",
	}

	m := gogu.NewMemoizer[string, any](time.Second, time.Minute)

	expensiveOp := func() (*gogu.Item[any], error) {
		// We simulate here an extensive operation.
		time.Sleep(2 * time.Second)

		foo := gogu.FindByKey(sampleMap, func(key string) bool {
			return key == "foo"
		})
		m.Cache.Set("item", foo, gogu.DefaultExpiration)

		res, err := m.Cache.Get("item")
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	// At the first call key1 does not exists, so the simulated extensive operation will take more time.
	data, err := m.Memoize("key1", expensiveOp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	// Second time the date will be read from cache. It will return the result instantly.
	data, err = m.Memoize("key1", expensiveOp)
	fmt.Println(data)
}
