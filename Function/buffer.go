package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.name))
	return err
}
func main() {
	p := person{name: "Bhumit Jain"}

	f, err := os.Create("output2.txt")
	if err != nil {
		log.Println(err)

	}
	defer f.Close()

	var b bytes.Buffer
	fmt.Printf("The type of b is %T\n", b)
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())

}
