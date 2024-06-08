package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(len(os.Args), os.Args)
		fmt.Println("Please enter a file to parse")
		return
	}
	filename := os.Args[1]
	fmt.Println("hi", filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Content of the file is wrong", err)
		return
	}

}

// read a file => os package
// parse
// if valid json-> 0
// if invalid json -> 1 - 255
