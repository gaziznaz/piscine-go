package piscine

func MakeRange(min, max int) []int {
	if min < max {

		size := max - min
		a := make([]int, size)
		for i := 0; i < size; i++ {
			a[i] = min + i
		}
		return a
	} else {
		return []int(nil)
	}
}
