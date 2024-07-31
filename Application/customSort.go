package main

import (
	"fmt"
	"sort"
)

type mySort []int

func (a mySort) Len() int           { return len(a) }
func (a mySort) Less(i, j int) bool { return a[i] < a[j] }
func (a mySort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func main() {
	arr := []int{67, 54, 23, 80, 89, 54}
	// or we can directly assign arr in below but need to initialize and declare with var arr mySort={...}
	sort.Sort(mySort(arr))
	fmt.Println(arr)

}
