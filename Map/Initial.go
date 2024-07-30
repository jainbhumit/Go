package main

import "fmt"

func main(){

	age:=map[string]int{
		"Bhumit":20,
		"Manvi":21,
		"Bhuvnesh":21,
	}
	fmt.Printf("the Age of Bhumit is %d",age["Bhumit"])
	var mymap map[string] int
	fmt.Println(mymap)

	var mymap2=make(map[string]int)
	mymap2["harsh"]=20
	mymap2["khushi"]=21

	fmt.Println(mymap2)

	for i,j:=range mymap2{
		fmt.Println("The age of ",i,"is ",j)
	}

	//deletee 
	delete(age,"Manvi")
	fmt.Println(age)
	//comma ok idiom 
	// v,ok:=age["Manvi"]

	// if ok {
	// 	fmt.Println("Present ",v)
	// }else{
	// 	fmt.Println("Didn't exist")
	// }

	//another way 
	

	if v,ok:=age["Manvi"];ok {
		fmt.Println("Present ",v)
	}else{
		fmt.Println("Didn't exist")
	}


}