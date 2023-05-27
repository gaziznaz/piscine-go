package piscine

func ToLower(s string) string {
	srune := []rune(s)
	var a rune

	for i := range s {
		if srune[i] >= 'A' && srune[i] <= 'Z' {
			a = rune(s[i] + 32)
			srune[i] = a
		}
	}
	return string(srune)
}
