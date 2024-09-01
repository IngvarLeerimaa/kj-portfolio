/*
Instructions
Write a function that receives a string slice and prints each element of
the slice in a seperate line.

Expected function
func PrintWordsTables(a []string) {

}
Usage
Here is a possible program to test your function :

package main

import "piscine"

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
}
And its output :

$ go run .
Hello
how
are
you?
$
*/
package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, x := range a {
		y := x
		for _, z := range y {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}
