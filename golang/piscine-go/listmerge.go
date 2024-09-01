package piscine

func ListMerge(l1 *List, l2 *List) {
	pos := l2.Head
	for pos != nil {
		ListPushBack(l1, pos.Data)
		pos = pos.Next
	}
}
