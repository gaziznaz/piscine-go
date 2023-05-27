package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	b := 0
	c := 0
	for i := range a {
		b = i + 1
	}
	for j := 0; j < n; j++ {
		c = j
	}

	if n <= 0 {
		return '\x00'
	} else if n > b {
		return 0
	}
	return a[c]
}
