package piscine

func ListPushFront(l *List, data interface{}) {
	if l.Head == nil {
		l.Head, l.Tail = &NodeL{Data: data}, l.Head
	} else {
		x := &NodeL{Data: data}
		x.Next, l.Head = l.Head, x
	}
}
