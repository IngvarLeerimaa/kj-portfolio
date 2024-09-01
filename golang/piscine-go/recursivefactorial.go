package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 20 {
		return 0
	} else if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
