package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := Strings{"peach", "banana", "kiwi"}
	fmt.Println("IsSorted:", sort.IsSorted(fruits))
	fmt.Println("Before:", fruits)
	sort.Sort(fruits)
	fmt.Println("IsSorted:", sort.IsSorted(fruits))
	fmt.Println("After:", fruits)
}
