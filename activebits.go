package piscine

func ActiveBits(n int) int {
	count := 0
	for i := 0; i < 8; i++ {
		if n%2 == 1 {
			count++
		}
		n = n / 2
	}
	return count
}
