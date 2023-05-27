package piscine

func Atoi(s string) int {
	srune := []rune(s)
	num := 0
	t := 1
	for j, i := range srune {
		if i == '-' && j == 0 {
			t = -1
		} else if i == '+' && j == 0 {
			t = 1
		} else if i >= '0' && i <= '9' {
			num = num*10 + int(i-48)
		} else {
			num = 0
			break
		}
	}
	return num * t
}
