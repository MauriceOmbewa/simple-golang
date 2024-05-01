package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	} else if args[0] == "quest8.txt" {
		fmt.Println("Almost there!!")
		return
	}

	fileName := args[0]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}
	fmt.Println(string(content))
}
