package piscine

func ListReverse(l *List) {
	current := l.Head
	prev := l.Head
	next := l.Head
	prev = nil
	next = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Tail = l.Head
	l.Head = prev
}
