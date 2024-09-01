package main

//Write a program that prints the name of the program.
/*Example of output :

student/piscine/printprogramname$ go build main.go
student/piscine/printprogramnane$ ./main
main
student/piscine/printprogramname$ go build
student/piscine/printprogramname$ ./printprogramname | cat -e
printprogramname$
student/piscine/printprogramname$ go build -o Nessy
student/piscine/printprogramname$ ./Nessy
Nessy
student/piscine/printprogramname$ */

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	xString := os.Args
	for index, y := range xString[0] {
		if index == 0 && y == '.' {
			continue
		} else if index == 1 && y == '/' {
			continue
		}
		z01.PrintRune(y)
	}
	z01.PrintRune('\n')
}
