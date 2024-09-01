package main

import (
	// "fmt"

	"github.com/01-edu/z01"
)

type asd struct {
	x int
	y int
}

func setPoint(ptr *asd) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &asd{} // {} points struct

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printItOut(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printItOut(points.y)
	z01.PrintRune('\n')
}

// z01.PrintRune(rune(points.x) - '0')
func printItOut(r int) {
	var x []rune
	for r != 0 {
		x = append(x, rune(r%10)+'0')
		r /= 10
	}
	for i := len(x) - 1; i >= 0; i-- {
		z01.PrintRune(x[i])
	}
}

// x := '0'
// if r == 0{
// 	z01.PrintRune('0')
// }
// if r > 0 {
// 	for i := 0; i < r
// }

// fmt.Printf("x = %d, y = %d\n",points.x, points.y)
