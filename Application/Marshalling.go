package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "Bhumit",
		Age:  20,
	}
	p2 := person{
		Name: "Bill",
		Age:  30,
	}

	candidate := []person{p1, p2}
	jForm, err := json.Marshal(candidate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jForm))

}
