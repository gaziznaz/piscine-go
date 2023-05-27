package piscine

func IsNumeric(s string) bool {
	for i := range s {
		if !(s[i] >= 48 && s[i] <= 57) {
			return false
		}
	}
	return true
}
