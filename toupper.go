package piscine

func ToUpper(s string) string {
	srune := []rune(s)
	var a rune

	for i := range s {
		if srune[i] >= 'a' && srune[i] <= 'z' {
			a = rune(s[i] - 32)
			srune[i] = a
		}
	}
	return string(srune)
}
