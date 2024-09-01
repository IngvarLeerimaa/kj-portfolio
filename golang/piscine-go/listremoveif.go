package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var slice []interface{}
	pos := l.Head
	for pos != nil {
		if pos.Data != data_ref {
			slice = append(slice, pos.Data)
		}
		pos = pos.Next
	}
	*l = List{}

	for i := 0; i < len(slice); i++ {
		ListPushBack(l, slice[i])
	}
}
