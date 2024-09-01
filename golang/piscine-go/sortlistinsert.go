package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listAdd(l, data_ref)
	var ints []int
	pos := l
	for pos != nil {
		ints = append(ints, pos.Data)
		pos = pos.Next
	}
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	posRes := l
	for i := 0; i < len(ints); i++ {
		posRes.Data = ints[i]
		posRes = posRes.Next
	}
	return l
}

func listAdd(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
