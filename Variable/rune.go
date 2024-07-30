package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "23456"
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\n", runes[i])

	}

	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d\n", number)
	}
	var char rune = '😁'

	fmt.Println(char)
}
