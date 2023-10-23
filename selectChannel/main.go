package main

import (
	"fmt"
	"time"
)

func main() {

	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		one <- "Hey"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		two <- "Hello"
	}()

	select {
	case rec1 := <-one:
		fmt.Println("I received from channel 1 : ", rec1)
	case rec2 := <-two:
		fmt.Println("I received from channel 2 : ", rec2)
	}

}

/// Tutorial
/// https://youtu.be/1c7ttSJDMAI
