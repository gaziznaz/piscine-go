package piscine

func BasicAtoi(s string) int {
	srune := []rune(s)
	num := 0
	for _, i := range srune {
		num = num*10 + int(i-48)
	}
	return num
}
