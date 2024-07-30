package main

import (
	"fmt"
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {

	return a + b
}
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addT(5, 6))
	fmt.Println(addT(5.6, 6.7))

}
