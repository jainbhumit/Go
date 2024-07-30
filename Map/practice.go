package main

import "fmt"
func main(){
	market:=map[string][]string{
		"Vegetable":{"potato","tomato","cauliflower"},
		"Fruit":{"mango","banana","gauava","apple"},
		"Grocery":{"rice","sugar","oil"},
	}

	fmt.Println(market)
	
	for k,v:=range market{
		fmt.Println("The item in ",k,"is : ")
		for _,k:=range v{
			fmt.Printf("%v ",k)
		}
		fmt.Printf("\n")
	}


}