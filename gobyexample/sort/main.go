package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	fmt.Println("Strings sorted:", sort.StringsAreSorted(strs))
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("Strings sorted:", sort.StringsAreSorted(strs))

	ints := []int{4, 6, 7, 1, 3, 2}
	fmt.Println("Ints sorted:", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	fmt.Println("Ints sorted:", sort.IntsAreSorted(ints))
}
