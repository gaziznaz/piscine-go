package piscine

func TrimAtoi(s string) int {
	res := 0
	n := 1

	for _, i := range s {
		if i >= '0' && i <= '9' {
			a := int(i - 48)
			res = res*10 + a
		} else if i == '-' && res == 0 {
			n = -1
		}
	}
	return res * n
}
