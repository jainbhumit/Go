package main

import "fmt"

type person struct{
	firstName string
	lastName string
	age int

}

func main(){
	p1:=struct{
		grade string 
		college string 
		
	}{
		grade :"A",
		college:"JECRC",
	}

	p2:=person{
		firstName:"Bhumit",
		lastName:"Jain",
		age:20,
	}


	fmt.Println(p1,p2)
	fmt.Printf("The p3 type is %T and val is : %v",p2,p2)
}