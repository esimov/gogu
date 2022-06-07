package main

import (
	"fmt"

	"github.com/esimov/gogu"
)

func main() {
	ints := []int{12, 23, 1, 643, 99, 2, 2}

	fmt.Println("==================FindAll")
	res1 := gogu.FindAll(ints, func(v int) bool {
		return v == 2
	})
	fmt.Println(res1)

	fmt.Println("==================FindAllFromLast")
	res2 := gogu.FindAllFromLast(ints, func(v int) bool {
		return v == 12
	})
	fmt.Println(res2)

	fmt.Println("==================FindIndex")
	res3 := gogu.FindIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res3)

	fmt.Println("==================FindLastIndex")
	res4 := gogu.FindLastIndex(ints, func(v int) bool {
		return v == 23
	})
	fmt.Println(res4)

	fmt.Println("==================IndexOf")
	fmt.Println(gogu.IndexOf(ints, 12))

	fmt.Println("==================LastIndexOf")
	fmt.Println(gogu.LastIndexOf(ints, 99))
}
