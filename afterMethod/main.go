package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//var wg sync.WaitGroup
	//wg.Add(1)
	go sell(ch)
	//time.Sleep(1 * time.Second)
	select {
	case sec := <-ch:
		fmt.Println("Received from channel", sec)
	case <-time.After(2 * time.Second):
		fmt.Println("Executed after 2 seconds")
		// default:
		// 	fmt.Println("default case executed")
	}
	//wg.Wait()
}

func sell(ch chan int) {
	//time.Sleep(5 * time.Second)
	ch <- 100 // Remove this line, as we're now sending in the main function.

	//wg.Done()
}
