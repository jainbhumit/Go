package main

import (
	"fmt"
)

type myType interface {
	~int | ~float64
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
type myalias int

func main() {
	fmt.Println(addT(5, 8))
	fmt.Println(addT(5.6, 6.7))

	// by using ~ in interface before the type helps to identify all the same type alias

	var x myalias = 20
	fmt.Println(addT(x, 10))

}
