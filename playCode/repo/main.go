package main

import (
	"fmt"
	testpackage "repo/testPackage"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(testpackage.Greet())
}
