package piscine

func BasicAtoi(s string) int {
	x := 0
	y := 0
	z := []rune(s)
	for _, word := range z {
		for i := '0'; i < word; i++ {
			y++
		}
		x = x*10 + y
		y = 0
	}
	return x
}
