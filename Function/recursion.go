package main

import(
"fmt"
)

func main(){
for i:=1;i<=5;i++{
	fmt.Println("The factorial of ",i," is : ",fac(i))
}
}

func fac(n int ) int {
	if n==0 || n==1 {
		return 1
	}
	return n*fac(n-1)

}