package piscine

func ConcatParams(args []string) string {
	s := []string(args)
	var str string
	len := 0

	for i := range s {
		len = i
	}
	for i := 0; i <= len; i++ {
		str = str + s[i]
		if i < len {
			str = str + "\n"
		}

	}
	return str
}
