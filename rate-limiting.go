package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i * 10
	}
	close(requests)

	limiter := time.Tick(100 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	fmt.Println()


	burstyLimiter := make(chan struct{}, 3)
	for range 3 {
		burstyLimiter <- struct{}{}
	}

	go func() {
		for range time.Tick(100 * time.Millisecond) {
			burstyLimiter <- struct{}{}
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		burstyRequests <- 100 * i
	}
	close(burstyRequests)

	for req :=  range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}