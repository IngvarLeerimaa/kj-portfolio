package piscine

func AlphaCount(s string) int {
	x := []rune(s)
	i := 0
	for y := range x {
		if (x[y] >= 'a' && x[y] <= 'z') || (x[y] >= 'A' && x[y] <= 'Z') {
			i++
		}
	}
	return i
}
