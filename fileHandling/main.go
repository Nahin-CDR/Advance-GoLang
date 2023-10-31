package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, []any{"\n----File Handling Program-----\n\n"}...)

	fileInfo, err := os.Stat("/Users/nazmulhaquenahin/temp.txt")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File Name : ", fileInfo.Name())
	fmt.Println("File Size : ", fileInfo.Size())
	fmt.Println("File Mode : ", fileInfo.Mode())
	fmt.Println("File IsDir : ", fileInfo.IsDir())

	println("\n----------------------------------\n")

	path := "/Users/nazmulhaquenahin/temp.txt"

	data, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}
	println("ASCII code of text ")
	fmt.Println(data)

	println("Real file : ", string(data))

	print("\n---------------------------------------------------\n")
}
