package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	res := false
	count := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			count++
		}
	}
	if count == len(a)-1 || count == 0 {
		res = true
	}
	return res
}
