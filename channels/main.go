package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sell(ch)
	go buy(ch)
	time.Sleep(1 * time.Second)
}

func sell(ch chan string) {
	ch <- "Furniture"
	fmt.Println("Sent data to the channel")
}

func buy(ch chan string) {
	fmt.Println("waiting for data")
	value := <-ch
	fmt.Println("data received : ", value)
}
