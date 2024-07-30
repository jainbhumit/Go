package main

import "fmt"

func call() {
	fmt.Println("This is function with no value and argument ")

}
func show(s string) {
	fmt.Println("My name is : ", s)
}

func give(s string) string {
	return fmt.Sprintln("The name is given by ", s)
}
func main() {
	//function with no parameter and no argument
	call()
	//function with paramee=te
	show("Bhumit")
	//with argument and return type
	s := give("Bhumit")
	fmt.Println(s)

}
