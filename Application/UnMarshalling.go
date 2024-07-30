package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"Name"`
	Last string `json:"Last"`
	Age  int    `json:"Age"`
}

func main() {
	// unmarshalling
	s := `[{"Name":"Bhumit","Last":"Jain","Age":20},{"Name":"Bill","Last":"Joe","Age":30}]`
	//convert into []byte
	bs := []byte(s)

	var people []person
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}
