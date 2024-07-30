package main

import "fmt"

func change(a []int) []int {
b:=make([]int,len(a))
copy(b,a)
for i:=0;i<len(a);i++{
	b[i]++
}
return b
}


func main(){
a:=[]int{5,6,7,8,9,10}
b:=a
fmt.Println(a)
fmt.Println(b)

a[0]=7
fmt.Println(a)
fmt.Println(b)
// to handle this situation we use copy function 

x:=[]int{56,78,90,76,54}

y:=make([]int,5)
copy(y,x)
x[0]=5678
fmt.Println(x)

fmt.Println(y)

fmt.Println(change(x))
fmt.Println(x)





}