package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	rep := []rune{}
	var ans string
	var ln int
	IsF := true

	for _, value := range arg {
		for _, v2 := range value {
			if check(v2) {
				rep = append(rep, v2)
				ln++
			}
		}
		if IsF {
			ans = value
			IsF = false
			continue
		}
		ans = ans + " " + value
	}
	cur := 0
	for index, value3 := range ans {
		if check(value3) {
			z01.PrintRune(rep[ln-cur-1])
			cur++
		} else {
			z01.PrintRune(value3)
		}
		if index == len(ans)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}
