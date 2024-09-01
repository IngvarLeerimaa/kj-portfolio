/*
Instructions
Write a program that prints the corresponding letter in the n position of the latin alphabet, where n is each argument received.

For example 1 matches a, 2 matches b, etc. If n does not match a valid position of the alphabet or if the argument is not an integer, the program should print a space (" ").

A flag --upper should be implemented. When used, the program prints the result in upper case. The flag will always be the first argument.

Usage
$ go run .
$ go run . 8 5 12 12 15 | cat -e
hello$
$ go run . 12 5 7 5 14 56 4 1 18 25 | cat -e
legen dary$
$ go run . 32 86 h | cat -e
   $
$ go run . --upper 8 5 25
HEY$
$
*/
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		arguments := os.Args[1:]
		suurt2ht := false
		if arguments[0] == "--upper" {
			suurt2ht = true
			arguments = os.Args[2:]
		}

		for _, z := range arguments {
			num := 0
			for _, x := range z {
				num = num*10 + int(x-'0')
			}
			if num >= 1 && num <= 26 {
				if !suurt2ht {
					z01.PrintRune(rune(num + 96))
				} else {
					z01.PrintRune(rune(num + 64))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
