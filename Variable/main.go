package main

import "fmt"

func main(){
	var sum int // initaial the value of Sum to 0 
	fmt.Println(sum)

	// var ch char='g'
	// fmt.Println(ch) //char is not used in the Go 

	sum=20
	fmt.Println(sum)

	a,b,c,d:="bhumit",5,6,4

	 fmt.Println(a,b,c,d)
	//fmt.Println(a,b,c) // this give error because variable d is unuse 
	fmt.Println("My name is ",a," Hello")

	x:=2
	y:='a'
	var z int  = x + int(y)
	fmt.Println(z)

}