package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	c := gogu.NewCache(5 * time.Minute)
	c.Set("foo", "bar", gogu.DefaultExpiration)

	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}
}
