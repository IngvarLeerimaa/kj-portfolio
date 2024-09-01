package piscine

func NRune(s string, n int) rune {
	x := []rune(s)
	if n > 0 && n <= len(s) {
		return x[n-1]
	}
	return 0
}
