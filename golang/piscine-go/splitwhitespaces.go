/*
Write a function that separates the words of a string and puts them in a string slice.

The separators are spaces, tabs and newlines.

Expected function
func SplitWhiteSpaces(s string) []string {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
}
And its output :

$ go run .
[]string{"Hello", "how", "are", "you?"}
$
*/
package piscine

func SplitWhiteSpaces(s string) []string {
	x := []rune(s)
	var wordsslice []string
	var word string

	for i := 0; i < len(s); i++ {
		if x[i] == ' ' || x[i] == '\n' || x[i] == '\t' {
			if word != "" {
				wordsslice = append(wordsslice, word)
				word = ""
			}
		} else {
			word = word + string(x[i])
		}
	}
	if word != "" {
		wordsslice = append(wordsslice, word)
	}
	return wordsslice
}
