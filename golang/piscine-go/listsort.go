package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
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
