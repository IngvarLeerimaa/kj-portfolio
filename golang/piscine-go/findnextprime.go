package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		var i int
		for i = nb + 1; !IsPrime(i); i++ {
		}
		return i
	}
}
