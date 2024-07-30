package main

import "fmt"

func main(){
	//creating a array 
	name:=[2]string{"Raj","Garvit"}
	//slice 
	// var arr[] string 
	arr1:=name[0:2]
	arr:= []string{"Bhumit","Raj","Garvit","Nishank"}
	fmt.Println(arr)

	for i,j:=range arr{
		fmt.Printf("The value at index %d is : %v\n",i,j)
	}
	arr[1]="Bhatt"

	arr= append(arr,"Babu")
	arr1= append(arr1,"Meet")
	arr= append(arr,arr1...)

	fmt.Println(arr)
	fmt.Println(arr1)

	//creating slice from the make function 
	highscore:=make([]int,4)
	highscore[0]=23
	highscore[1]=26
	highscore[2]=28
	highscore[3]=29
	// highscore[4]=200 throw error but go through append not throw eror
	highscore=append(highscore, 45,667,34)


	fmt.Println(highscore)

	//multidimentional slice 
	table:=[][]int{{5,6},{6,7},{8,10}}

	fmt.Printf("The length is %d",len(table[0]))



}