package main

import "fmt"

type Number interface {
	int | int32 | float64 | float32
}

func sum[N Number](a, b N) N {
	return a + b
}

func main() {
	res := sum(1.1, 2.2)
	fmt.Println(res)
}
