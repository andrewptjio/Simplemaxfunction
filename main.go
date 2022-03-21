package main

import (
	"edenfarm/functions/max"
	"fmt"
)

func main() {
	res := max.Max([]int{1, 2, 3, 8, 9, 3, 2, 1})
	fmt.Println(res)
	res = max.Max([]int{1, 2, 1, 2, 2, 1})
	fmt.Println(res)
	res = max.Max([]int{7, 1, 2, 9, 7, 2, 1})
	fmt.Println(res)
}
