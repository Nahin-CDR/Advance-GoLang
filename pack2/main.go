package main

import (
	"fmt"
	"github.com/Nahin-CDR/pack1/encrypt"
)

func main() {
	/// at first we will encrypt string
	name := encrypt.Nimbus("Nahin the Coder")
	fmt.Println("Welcome to Package 2")
	fmt.Println("Encoded String : ", name)

}
