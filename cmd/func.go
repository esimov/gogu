package main

import (
	"fmt"
	"time"

	"github.com/esimov/gogu"
)

func main() {
	fmt.Println("==================Delay")
	// gogu.Delay(time.Second, func() {
	// 	fmt.Println("finished")
	// })

	fmt.Println("==================After")
	sample := []int{1, 2, 3, 4, 5, 6}
	length := len(sample) - 1
	gogu.ForEach[int](sample, func(val int) {
		fmt.Printf("Printing value... %d\n", val)

		gogu.After[string, int](&length, func() {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("save after")
		})
	})

	fmt.Println("==================Before")
	var n int = 3
	gogu.ForEach[int](sample, func(val int) {
		res := gogu.Before[string, int](&n, func() string {
			time.Sleep(time.Millisecond * 100)
			return "memoized function"
		})
		fmt.Printf("Printing value... %d %v\n", val, res)
	})
}
