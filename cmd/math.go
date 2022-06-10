package main

import (
	"fmt"

	"github.com/esimov/gogu"
)

func main() {
	fmt.Println("==================Clamp")
	fmt.Println(gogu.Clamp(10, -5, 5))
}
