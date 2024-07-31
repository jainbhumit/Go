package main

import (
	"encoding/json"
	"fmt"
)

type intern struct {
	Name    string
	Batch   int
	College string
}

func main() {
	p1 := intern{"Bhumit", 4, "JECRC"}
	p2 := intern{"Gaurav", 4, "SKIT"}

	student := []intern{p1, p2}

	jForm, err := json.Marshal(student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jForm))

}
