package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	as := true
	de := true

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			as = false
		}
	}

	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) < 0 {
			de = false
		}
	}
	return as || de
}

// for index, i := range a {
// 	if index < len(a) {
// 		if i > a[i+1] {
// 			return false
// 		}
// 	}
// }
// return true
