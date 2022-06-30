package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	m := gogu.NewMemoizer[string, any](time.Second, time.Minute)

	expensiveOperation := func() (*gogu.Item[interface{}], error) {
		time.Sleep(2 * time.Second)
		bigDataStructure := struct{ key string }{key: "key1"}

		return &gogu.Item[interface{}]{
			Object: bigDataStructure,
		}, nil
	}

	data, err := m.Memoize("key1", expensiveOperation)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	// Return data instantly
	data, err = m.Memoize("key1", expensiveOperation)
	fmt.Println(data)
}
