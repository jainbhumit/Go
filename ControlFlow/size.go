package main

import "fmt"
import "unsafe"

func main()  {
	a,b,c :=10,"Bhumit",5.6

	fmt.Printf("Size of %T is : %d \n",a,unsafe.Sizeof(a))
	fmt.Printf("Size of %T is : %d \n",b,unsafe.Sizeof(b))
	fmt.Printf("Size of %T is : %d \n",c,unsafe.Sizeof(c))

}