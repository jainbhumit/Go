package main

import "fmt"

func main() {
	//arr := []interface{}{"Bhumit", 5, 7.9, true}
	//fmt.Println(arr)
	//OR
	var arr []interface{}
	arr = append(arr, "A", 56, 7.8, true)

	for _, v := range arr {
		fmt.Printf("The type of %v is : %T\n", v, v)
		
	}
}
