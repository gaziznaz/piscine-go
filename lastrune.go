package piscine

func LastRune(s string) rune {
	a := []rune(s)
	c := 0
	// count := 0
	for index := range a {
		c = index
	}
	return a[c]
}
