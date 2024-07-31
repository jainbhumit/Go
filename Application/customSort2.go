package main

import (
	"fmt"
	"sort"
)

type people []person
type person struct {
	name string
	age  int
}

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i].age > p[j].age }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func main() {
	var member people = people{
		{"Bhumit", 20},
		{"Garvit", 21},
		{"Babu", 50},
		{"Naman", 19},
	}
	sort.Sort(member)
	fmt.Println(member)

}
