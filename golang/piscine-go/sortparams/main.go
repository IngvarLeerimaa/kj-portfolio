/*
Instructions
Write a program that prints the arguments received in the command line in ASCII order.

Example of output :

$ go run . 1 a 2 A 3 b 4 C
1
2
3
4
A
C
a
b
$
*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	x := []string(os.Args[1:])
	for j := 0; j < len(x); j++ {
		for i := 0; i < len(x)-1; i++ {
			if x[i] > x[i+1] {
				a := x[i]
				b := x[i+1]
				x[i] = b
				x[i+1] = a
			}
		}
	}
	for y := 0; y < len(x); y++ {
		for _, w := range x[y] {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}
