package main

import (
	"fmt"
	"time"
)

func start() {
	go inProgress()
	fmt.Println("start")
}

func inProgress() {

	fmt.Println("inprogress")
}

func main() {
	start()

	time.Sleep(1 * time.Second)
	fmt.Println("Execution ended !")
}
