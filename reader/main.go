package main

import (
	"fmt"
	"strings"
)

func main() {

	r := strings.NewReader("Learning is Fun")

	buf := make([]byte, 4)

	for {
		n, err := r.Read(buf)
		fmt.Println(buf[:n], err)
		if err != nil {
			fmt.Println("Breaking out with error : ", err)
			break
		}
	}

}
