package main

import (
	"fmt"
	"unsafe"
)

func change(a *int) {
	*a++
}

func print(arr []int) {
	arr[0] = 9
	fmt.Println(arr)
}
func main() {
	a := 10
	fmt.Println(a)
	change(&a)
	fmt.Println(a)

	arr := [5]int8{1, 2, 3, 4, 5}

	for i, _ := range arr {
		fmt.Println(&arr[i], " ", unsafe.Sizeof(arr))
		fmt.Printf("%T\n", arr[i])
	}

	//print(arr)
	//fmt.Println(arr)
	exp, err := fmt.Println("The example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(exp)

}
