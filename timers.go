package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 done waiting")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 done waiting")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}