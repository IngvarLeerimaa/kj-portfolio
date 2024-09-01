package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	pos := l.Head
	for pos != nil {
		if comp(ref, pos.Data) {
			return &pos.Data
		}
		pos = pos.Next
	}
	return nil
}
