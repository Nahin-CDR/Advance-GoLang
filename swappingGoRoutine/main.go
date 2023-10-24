package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 1; i <= 10; i++ {
		go func(k int) {
			fmt.Println(k)
		}(i)
		wg.Done()
	}
	fmt.Println("Program Ended")
	wg.Wait()
}
