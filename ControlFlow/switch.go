package main
import "fmt"

func main(){
	// var x int

	// fmt.Printf("Enter the no : \n")
	
	// fmt.Scanln(&x)

	// switch{
	// case x>50: 
	// 	fmt.Println("Greater")
	// case x==50:
	// 	fmt.Println("Equal")
	// case x<50:
	// 	fmt.Println("Small")
	// default:
	// 	fmt.Println("Invalid")


	// }

	var grade string 
	fmt.Printf("Enter the Grade : \n")
	fmt.Scanln(&grade)

	switch grade{
	case "A": 
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Best")
	case "C":
		fmt.Println("good")
	default:
		fmt.Println("Fail")


	}
}