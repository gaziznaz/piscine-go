package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		return
	} else {
		n.Next = l.Head
		l.Head = n

	}
}
