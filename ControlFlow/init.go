package main

import "fmt"

var a,b int=5,6
func init(){
fmt.Println("This is initializer function ")
a,b=7,8
}
func main()  {
	fmt.Print("This is main function and the value of a and b is :  ",a,b)

}