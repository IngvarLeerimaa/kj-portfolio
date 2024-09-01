package main

/*
WWrite a program that prints the arguments received in the command line in reverse order.

Example of output :

$ go run . choumi is the best cat
cat
best
the
is
choumi
$
*/

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	xString := os.Args
	ln := 0 // counter
	for range xString {
		ln++
	}
	for y := ln - 1; y > 0; y-- {
		for _, w := range xString[y] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
