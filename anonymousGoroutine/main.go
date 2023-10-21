package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("In anonymous method")
	}()
	time.Sleep(2 * time.Second)
}
