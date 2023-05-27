package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	m := l
	size := 0
	for m != nil {
		m = m.Next
		size++
	}
	if l == nil {
		return nil
	} else if size < pos {
		return nil
	} else {
		n := l

		for i := 0; i < pos; i++ {
			n = n.Next
		}
		return n
	}
}
