package main

import "fmt"

func main() {
	//syntax
	func() {
		fmt.Println("Anonymous function ")
	}() //used for immediate invoke of function

	//storing in another variable by passing parameter
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(5, 6))

}
