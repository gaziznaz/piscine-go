package piscine

func Rot14(s string) string {
	srune := []rune(s)
	var res []rune
	for _, k := range srune {
		if (k >= 65 && k <= 76) || (k >= 97 && k <= 108) {
			res = append(res, k+14)
		} else if (k >= 77 && k <= 90) || (k >= 109 && k <= 122) {
			res = append(res, k-12)
		} else {
			res = append(res, k)
		}
	}
	return string(res)
}
