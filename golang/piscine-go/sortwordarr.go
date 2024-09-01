package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				swap(&a[i], &a[j])
			}
		}
	}
}

func swap(a, b *string) {
	x := *a
	y := *b
	*a = y
	*b = x
}
