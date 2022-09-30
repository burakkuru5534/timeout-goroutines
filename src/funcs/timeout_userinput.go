package funcs

import (
	"fmt"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	c := make(chan int)
	go func() {
		//defer close(c)
		time.Sleep(5 * time.Second)
		w.Wait()
		c <- 1
	}()
	select {
	case <-c:
		return false
	case <-time.After(t):
		return true
	}
}

func TimeOutByUserInput() {
	var t int32
	fmt.Printf("Enter duration for timeout(in milisecond):")
	fmt.Scanf("%d", &t)

	var w sync.WaitGroup
	w.Add(1)
	d := time.Duration(t) * time.Millisecond
	fmt.Printf("\n...will timeout after %s ...\n", d)
	if timeout(&w, d) {
		fmt.Println("TIMEOUT!")
	} else {
		fmt.Println("DONE!")
	}
	w.Done()
	if timeout(&w, d) {
		fmt.Println("TIMEOUT!")
	} else {
		fmt.Println("DONE!")
	}
}
