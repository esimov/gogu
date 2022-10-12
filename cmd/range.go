package main

import (
	"fmt"

	"github.com/esimov/torx"
)

func main() {
	fmt.Println("==================Range")
	if rn, err := torx.Range[int](-10, -2, -320); err == nil {
		fmt.Println(rn)
	}

	fmt.Println("==================Range floats")
	rf, _ := torx.Range[float64](0.0, 0.1, 1)
	fmt.Println(rf)

	fmt.Println("==================RangeRight")
	if rr, err := torx.RangeRight[int](0, -1, -4); err == nil {
		fmt.Println(rr)
	}
}
