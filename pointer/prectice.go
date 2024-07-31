package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{"Bhumit"}
	fmt.Println(p)

	newP := change(p, "Garvit")
	fmt.Println(newP)
	fmt.Println(p)

	changePtr(&p, "Babu")
	fmt.Println(p)

}
func change(p person, s string) person {
	p.name = s
	return p
}

func changePtr(p *person, s string) {
	p.name = s
}
