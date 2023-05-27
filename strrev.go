package piscine

func StrRev(s string) string {
	srune := []rune(s)
	var runerev []rune
	len1 := len(srune) - 1

	for i := len1; i >= 0; i-- {
		runerev = append(runerev, srune[i])
	}
	return string(runerev)
}
