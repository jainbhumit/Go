package main

import(
	"fmt"
)

func main(){
defer fmt.Println(4)
defer fmt.Println(3)
defer fmt.Println(2)
defer fmt.Println(1)
fmt.Println(0)
}

//defer act as a LIFO property 