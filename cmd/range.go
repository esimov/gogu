package main

import (
	"fmt"

	"github.com/esimov/gogu"
)

func main() {
	fmt.Println("==================Range")
	rn, err := gogu.Range[int](-10, -2, -320)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rn)
	}

	fmt.Println("==================RangeRight")
	rr, err := gogu.RangeRight[int](0, -1, -4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rr)
	}
}
