package main

import "fmt"
func main(){

	a:=5
	fmt.Println(&a)
}

// to check the static and heap use go run -gcflags -m filename.go
