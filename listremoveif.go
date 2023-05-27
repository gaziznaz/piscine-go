package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	new := &NodeL{}
	current := l.Head
	temp := new

	for current != nil {

		if current.Data != data_ref {
			temp.Next = current
			temp = temp.Next
		}
		current = current.Next
	}
	temp.Next = nil
	l.Head = new.Next
}
