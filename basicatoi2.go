package piscine

func BasicAtoi2(s string) int {
	srune := []rune(s)
	num := 0
	for _, i := range srune {
		if i >= 48 && i <= 57 {
			num = num*10 + int(i-48)
		} else {
			num = 0
			break
		}
	}
	return num
}
