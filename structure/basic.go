package main

import "fmt"
type person struct{
	firstName string
	lastName string
	age int

}

type address struct{
	person
	add string
}


func main(){
	p1:=person{
		firstName:"Bhumit",
		lastName:"Jain",
		age:20,
	}
	p2:=person{
		firstName:"Garvit",
		lastName:"Jain",
		age:21,
	}

	// fmt.Println(person)
	fmt.Println(p1)
	fmt.Println(p2)

	personAddress:=address{
		person:p1,
		add:"Bahubali,Banswara",

	}
	fmt.Println(personAddress)
}
