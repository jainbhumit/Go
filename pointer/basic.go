package main

import "fmt"

func main() {
	//var x int = 32
	//p := &x
	//fmt.Println(x)
	//fmt.Println(&x)
	//fmt.Println(p)
	//fmt.Println(*p)
	//fmt.Printf("%T\n", x)
	//fmt.Printf("%T\n", p)
	//fmt.Printf("%T\n", &x)
	//fmt.Printf("%T\n", *p)
	//
	//var ptr *int
	//ptr = &x
	//
	//fmt.Printf("The type is %T", ptr)
	//
	//hands on On the pointer
	a := 5
	ptr1 := &a
	ptr2 := &ptr1
	fmt.Println(a)      //5
	fmt.Println(&a)     // address of a
	fmt.Println(ptr1)   //adress of a
	fmt.Println(*ptr1)  // 5
	fmt.Println(&ptr1)  //adress of ptr1
	fmt.Println(*ptr2)  //adress of a
	fmt.Println(ptr2)   //adress of ptr1
	fmt.Println(&ptr2)  // adress of ptr2
	fmt.Println(**ptr2) // 5

}
