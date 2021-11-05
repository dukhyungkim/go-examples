package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}
