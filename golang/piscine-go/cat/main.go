package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func PrintRune(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
}

func main() {
	size := len(os.Args)

	for i := 1; i < size; i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			PrintRune("ERROR: ")
			PrintRune(err.Error())
			z01.PrintRune('\n')
			os.Exit(1)
			return
		}
		PrintRune(string(data))
		// PrintRune()
	}
	bytes, _ := ioutil.ReadAll(os.Stdin)
	PrintRune(string(bytes))
}
