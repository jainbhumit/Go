package main

import (
	"fmt"
	"log"
	"strconv"
)

type person struct {
	name string
}

func (p person) String() string {
	return fmt.Sprintf("This is person and name is %v", p.name)

}

type count int

func (c count) String() string {
	return fmt.Sprintf("This is count %v", strconv.Itoa(int(c)))

}

func logPrint(obj fmt.Stringer) {
	fmt.Println("This is wrapper and belongs to ", obj.String())
}
func main() {
	c := count(6)
	p := person{name: "Bhumit"}
	fmt.Println(p)
	fmt.Println(c)
	log.Println(p)
	log.Println(c)
	logPrint(p)
	logPrint(c)

}
