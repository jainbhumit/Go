package main

import "fmt"

type Animal struct {
	name string
}
type Dog struct {
	Animal
}
type Cat struct {
	Animal
}

func (obj Dog) Speak() {
	fmt.Println("Dog is Barking")

}
func (obj Cat) Speak() {
	fmt.Println("Cat is mewing")

}
func (obj Dog) Eat() {
	fmt.Println("Dog is Eating")

}
func (obj Cat) Eat() {
	fmt.Println("Cat is Eating")
}

type mammal interface {
	Speak()
	Eat()
}

func saySomething(m mammal) {
	fmt.Println("This is mammal speaking and name is ")
	m.Speak()
	m.Eat()
}
func main() {
	gs := Dog{Animal{name: "Tommy"}}
	ki := Cat{Animal{name: "Kitty"}}
	gs.Speak()
	ki.Speak()
	gs.Eat()
	ki.Eat()
	saySomething(gs)
	saySomething(ki)

}
