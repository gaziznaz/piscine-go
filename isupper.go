package piscine

func IsUpper(s string) bool {
	str := []rune(s)

	bool1 := true
	for _, i := range str {
		if !(i >= 'A' && i <= 'Z') {
			bool1 = false
			return bool1
		}
	}
	return bool1
}
