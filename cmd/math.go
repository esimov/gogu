package main

import (
	"fmt"

	"github.com/esimov/torx"
)

func main() {
	fmt.Println("==================Clamp")
	fmt.Println(torx.Clamp(10, -5, 5))
}
