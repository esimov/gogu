package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	c := gogu.NewCache[string, string](1*time.Second, 1*time.Minute)
	err := c.Set("foo", "bar", gogu.DefaultExpiration)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)

	_, err = c.Get("foo")
	if err != nil {
		fmt.Println(err)
	}
	err = c.Set("foo", "bar", gogu.DefaultExpiration)
	if err != nil {
		fmt.Println(err)
	}

	items := c.List()
	for key, val := range items {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	// Key not found.
	fmt.Println(c.Get("test"))

	c.DeleteExpired()
}
