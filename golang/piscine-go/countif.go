package piscine

func CountIf(f func(string) bool, tab []string) int {
	x := 0
	for _, i := range tab {
		if f(i) {
			x++
		}
	}
	return x
}
