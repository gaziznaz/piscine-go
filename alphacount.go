package piscine

func AlphaCount(s string) int {
	sbyte := []byte(s)
	c := 0
	for i := range sbyte {
		if (sbyte[i] >= 65 && sbyte[i] <= 90) || (sbyte[i] >= 97 && sbyte[i] <= 122) {
			c++
		}
	}
	return c
}
