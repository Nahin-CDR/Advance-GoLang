package main

import (
	"fmt"

	"github.com/Nahin-CDR/pack1/decrypt"
	"github.com/Nahin-CDR/pack1/encrypt"
)

func main() {
	/// at first we will encrypt string
	name := encrypt.Nimbus("Nahin the Coder")
	fmt.Println("Name after encrypt : ", name)
	fmt.Println("Name after decrypt : ")
	fmt.Println(decrypt.Nimbus(name))
}
