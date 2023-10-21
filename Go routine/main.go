package main

import (
	"fmt"
	"time"
)

func calculate(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	println("Welcome! Nahin the coder")
	startTime := time.Now()
	for i := 1; i <= 10; i++ {
		go calculate(i)
	}

	elapsed := time.Since(startTime)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took time : ", elapsed)

}
