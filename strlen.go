package piscine

func StrLen(s string) int {
	srune := []rune(s)
	count := 0
	for i := range srune {
		count = i + 1
	}
	return count
}
