package piscine

func ListMerge(l1 *List, l2 *List) {
	first := l1.Head

	if l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	} else if l2.Head == nil {
		return
	}
	for first.Next != nil {
		first = first.Next
	}
	first.Next = l2.Head
	l1.Tail = l2.Tail
}
