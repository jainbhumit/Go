package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)

	}
	defer f.Close()

	text := []byte("This is writing learning in file")
	fmt.Printf("the type of text is %T\n", text)
	_, err = f.Write(text)
	if err != nil {
		log.Fatalf("error %s", err)
	}

}
