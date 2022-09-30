package funcs

import (
	"fmt"
	"time"
)

/*
Examples of How to timeout Out a Goroutine in Go: Using time.Sleep()
*/

func TimeOutByUsingTimeSleep() {

	channel1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		channel1 <- "CHANNEL 1"
	}()

	select {
	case str := <-channel1:
		fmt.Println(str)
	case <-time.After(time.Second * 1):
		fmt.Println("TIMEOUT: CHANNEL 1")
	}

	channel2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 4)
		channel2 <- "CHANNEL 2"
	}()

	select {
	case str := <-channel2:
		fmt.Println(str)
	case <-time.After(time.Second * 3):
		fmt.Println("TIMEOUT: CHANNEL 2")
	}
}
