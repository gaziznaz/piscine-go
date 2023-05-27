package piscine

func ListSize(l *List) int {
	size := 0
	i := l.Head
	for i != nil {
		i = i.Next
		size++
	}
	return size
}
