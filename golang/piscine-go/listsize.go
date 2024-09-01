package piscine

func ListSize(l *List) int {
	counter := 0
	x := l.Head
	for x != nil {
		counter++
		x = x.Next
	}
	return counter
}
