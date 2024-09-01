package main

/*
Instructions
Write a program that prints the arguments received in the command line.

Example of output :

$ go run . choumi is the best cat
choumi
is
the
best
cat
$
*/

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	xString := os.Args
	ln := 0 // counter
	for y := range xString {
		ln = y
	}
	for y := 1; y <= ln; y++ {
		for _, w := range xString[y] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
