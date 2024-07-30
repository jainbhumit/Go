package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type myType interface {
	constraints.Integer | constraints.Float
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {

	return a + b
}
func addT[T myType](a, b T) T {
	return a + b
}

// using alias ~

func main() {
	fmt.Println(addT(5, 8))
	fmt.Println(addT(5.6, 6.7))

}
