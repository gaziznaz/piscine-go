package piscine

func Unmatch(a []int) int {
	for _, b := range a {
		c := 0
		for _, b1 := range a {
			if b1 == b {
				c++
			}
		}
		if c == 1 || c%2 == 1 {
			return b
		}
	}
	return -1
}
