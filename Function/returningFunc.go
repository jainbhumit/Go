package main

import "fmt"

func dog() func() string {
	return func() string {
		return "Dog is Barking Woof Woof!"
	}
}
func main() {
	animal := dog()
	fmt.Println(animal())
}
