package piscine

func IsLower(s string) bool {
	str := []rune(s)

	bool1 := true
	for _, i := range str {
		if !(i >= 'a' && i <= 'z') {
			bool1 = false
			return bool1
		}
	}
	return bool1
}
