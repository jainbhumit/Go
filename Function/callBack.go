package main

import "fmt"

func main(){
adding:=math(13,4,add)
subtracting:=math(13,4,sub)

fmt.Println(adding)
fmt.Println(subtracting)


}


func math(a int,b int,do func(int,int)int) int{
	return do(a,b)
}

func add(a int,b int) int {
	return a+b
}

func sub(a int,b int) int {
	return a-b
}
