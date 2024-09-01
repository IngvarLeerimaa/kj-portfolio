package piscine

func BasicAtoi2(s string) int {
	x := 0
	y := 0
	z := []rune(s)
	for _, value := range z {
		if value < '0' || value > '9' {
			return 0
		}
		for i := '0'; i < value; i++ {
			y++
		}
		x = x*10 + y
		y = 0
	}
	return x
}
