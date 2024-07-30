package main

import "fmt"

type Car struct{
	Name string
	modelNo int

}

//creating the function for type car cause we don't apply function in the structure 
func (c Car)Run() {
	fmt.Println("Car is running")

} 

func main(){
c:=Car{
	Name:"BMW",
	modelNo:45,

}
c.Run()

}
