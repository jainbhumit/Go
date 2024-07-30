package main

import(
"fmt"
)

type Dog struct{
	name string
}

func (d Dog)Walk(){
	fmt.Println("The Dog name is : ",d.name," and its is walking.....")
}

func (d *Dog)Run(){
	d.name="Tommy"
	fmt.Println("The Dog name is : ",d.name," and its is running.....")

}

//here this code take the both sementic but in case of interface the pointer recieiver function used only by the object having the addres
//of the type for example here i create the interfae of run 
type activity interface{
	Run() // run is define as a reciver pointer 

}

func doActivity(a activity){
	a.Run()

}

func main(){
d1:=Dog{
	name :"Tom",
}
d2:=&Dog{"Tarzan"}

d1.Walk()
d1.Run()
d2.Walk()
d2.Run()
fmt.Println(d1.name)
fmt.Println(d2.name)

// now using the interface pattern 
//doActivity(d1) // it show the error 
doActivity(d2)



}