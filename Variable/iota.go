package main

import "fmt"
type byteSize int 

func main()  {
const (
	_ = iota// initaially 0 then increment 
	KB byteSize= 1<<(10*iota)
	MB 
	GB
	TB
	PB
	EB
)	

fmt.Println(KB,MB,GB,TB,PB,EB)
}