package piscine

func Any(f func(string) bool, a []string) bool {
	var res bool
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			res = true
			break
		}
		res = false
	}
	return res
}
