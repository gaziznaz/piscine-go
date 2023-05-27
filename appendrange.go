package piscine

func AppendRange(min, max int) []int {
	var a []int
	if min < max {
		for i := min; i < max; i++ {
			a = append(a, i)
		}
	}
	return a
}
