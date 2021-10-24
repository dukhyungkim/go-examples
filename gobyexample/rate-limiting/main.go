package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("SimpleLimiter")
	SimpleLimiter()

	fmt.Println()
	fmt.Println("BurstyLimiter")
	BurstyLimiter()
}

func SimpleLimiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

func BurstyLimiter() {
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
