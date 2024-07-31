package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Foo : ", i)
	}
	wg.Done()
}
func boo() {
	for i := 0; i < 5; i++ {
		fmt.Println("Boo : ", i)
	}
}
func main() {
	fmt.Println("OS    ", runtime.GOOS)
	fmt.Println("ARC    ", runtime.GOARCH)
	fmt.Println("CPU    ", runtime.NumCPU())
	fmt.Println("GORouting    ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	boo()
	fmt.Println("CPU    ", runtime.NumCPU())
	fmt.Println("GORouting    ", runtime.NumGoroutine())
	wg.Wait()
}
