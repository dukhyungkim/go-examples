package main

import (
	"fmt"
	"time"
)

// 일정 주기마다 무언가를 실행하기 위해선 sleep 보다 ticker 가 더 좋겠다.

func main() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker Stopped")
}
