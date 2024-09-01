/*Instructions
Write a function that returns a concatenated string from the 'strings' passed as arguments.

Expected function
func BasicJoin(elems []string) string {

}
Usage
Here is a possible program to test your function :

package main

import (
	"fmt"
	"piscine"
)

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(elems))
}
And its output :

$ go run .
Hello! How are you?
$
*/
package piscine

func BasicJoin(elems []string) string {
	var asd string
	for i := 0; i < len(elems); i++ {
		asd += string(elems[i])
	}
	return asd
}
