package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	c := gogu.NewCache[string, string](1 * time.Second)
	err := c.Set("foo", "bar", gogu.DefaultExpiration)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2 * time.Second)

	foo, err := c.Get("foo")
	if err != nil {
		fmt.Println(foo)
	}
	err = c.Set("foo", "bar", gogu.DefaultExpiration)
	if err != nil {
		fmt.Println(err)
	}
}
