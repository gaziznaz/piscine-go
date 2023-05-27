package piscine

func IsAlpha(s string) bool {
	for i := range s {
		if (s[i] >= 0 && s[i] < 48) || (s[i] > 57 && s[i] < 65) ||
			(s[i] > 90 && s[i] < 97) || (s[i] > 122 && s[i] <= 127) {
			return false
		}
	}
	return true
}
