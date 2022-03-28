package main

import (
	"fmt"

	"github.com/reactivex/rxgo/v2"
)

func main() {
	ch := make(chan rxgo.Item)
	go func() {
		for i := 0; i < 9; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	observable := rxgo.FromChannel(ch)

	// first observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}

	// second observer
	for item := range observable.Observe() {
		fmt.Println(item.V)
	}
}
