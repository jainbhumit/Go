package main

import (
	"fmt"
	math2 "math"
)

func main() {
	x := powinator(2)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

}
func powinator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math2.Pow(a, c)
	}

}
