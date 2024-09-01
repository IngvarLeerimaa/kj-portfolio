package piscine

func LastRune(s string) rune {
	x := []rune(s)
	i := len(s) - 1
	return x[i]
}
