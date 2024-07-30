package main

import "fmt"

func main(){
	x:=1

	//for loop like C 
	for i:=0;i<=10;i++{
	fmt.Printf("Value of i is %d \n",i)

	}

	// while loop like C
	for x<5{
		fmt.Println("The value of x is : ",x)
		x++;
	}

	//break
	for {
		fmt.Println("The value of x is : ",x)
		if x>10{
			break;
		}
		x++;
		
	}

	//coninue
	
	for x=0;x<10;x++{
		if x%2!=0{
			continue;
		}
		fmt.Println("The value of x is : ",x)
	
		
	}

	// loop in range range give the iterator index + value 

	arr:=[]int{5,6,7,8,9}

	for i,j:=range arr{
		fmt.Println("The value at index ",i," is ",j)
	}

}