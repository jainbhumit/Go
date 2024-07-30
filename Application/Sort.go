package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 20, 15, 32, 10, 14}

	si := []string{"Bhumit", "Garvit", "Bhuvnesh", "Harsh"}
	sort.Ints(xi)
	sort.Strings(si)
	fmt.Println(xi)
	fmt.Println(si)

}
