package main

import (
	"fmt"
	"strings"
)

func main() {
	num := 12
	name := "Arfan"
	// integer format specifier
	fmt.Printf("Hello world %d \n", num)
	// string format specifier
	fmt.Printf("Welcome %s to our office\n", name)

	learning := "Learning Standard Library in Go"
	fun := "Library in Go"

	result := strings.Contains(learning, fun)

	println(result)

}
