package piscine

func Map(f func(int) bool, a []int) []bool {
	x := []bool{}
	for _, i := range a {
		x = append(x, f(i))
	}
	return x
}
