package main

import "fmt"

func sum(ii ...int) int {
	s := 0
	for _, j := range ii {
		s += j
	}
	return s
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	//or
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(arr...))
}
