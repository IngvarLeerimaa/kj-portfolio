package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	}
	if len(os.Args) == 2 {
		file, _ := os.Open("quest8.txt")
		arr := make([]byte, 14)
		file.Read(arr)
		fmt.Println(string(arr))
	}
}
