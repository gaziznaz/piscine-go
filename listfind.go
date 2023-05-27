package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	n := l.Head
	var res interface{}
	for n != nil {
		res = n.Data
		if comp(ref, res) == true {
			return &res
		}
		n = n.Next

	}
	return &res
}
