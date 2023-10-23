package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan int)

	go sell(ch, &wg)
	go buy(ch, &wg)
	wg.Wait()
	simple()
}

func simple() {
	ch := make(chan int, 3)
	ch <- 200
	ch <- 900
	close(ch)
	println("Iterate this channel")
	for item := range ch {
		fmt.Println(item)
	}

}

func sell(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("10 injected")
	ch <- 10

	fmt.Println("20 injected")
	ch <- 20
	fmt.Println("30 injected")
	ch <- 30
	fmt.Println("40 injected")
	ch <- 40

	close(ch)
	wg.Done()
	fmt.Println("Sent all data")
}

func buy(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for data")

	for val := range ch {
		time.Sleep(2 * time.Second)
		fmt.Println("Received value", val)
	}
	wg.Done()
}
