package piscine

import (
	//	"fmt"

	"github.com/01-edu/z01"
)

/*Write a function which prints the digits of an int passed in parameter in ascending order.
All possible values of type int have to go through, excluding negative numbers. Conversion to int64 is not allowed.
*/

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0') // if n = 0; prints 0
		// return
	}
	var x [10]int // var x is 10 int array
	for n != 0 {  // if not 0
		x[n%10]++ // adds value to current array spot
		// fmt.Println(x)
		n /= 10 // var n int = n / 10 example n = 20/10 = 2
	}
	for i := 0; i < 10; i++ { // loop to read index
		for x[i] > 0 { // if value of index  its larger than 0 print it out
			z01.PrintRune(rune(i) + '0') // prints out index +'0' et ta aru saaks et tegemist on ascii numbriga
			x[i]--                       // post loop statement, to delete printed out current index value
		}
	}
}
