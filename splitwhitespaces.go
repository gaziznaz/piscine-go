package piscine

func SplitWhiteSpaces(s string) []string {
	srune := []rune(s)
	var word string
	var str []string

	i := 0

	for _, k := range srune {
		if word != "" && (k == ' ' || k == '	' || k == '\n') {

			str = append(str, word)
			word = ""
			i++

		} else {
			if k != ' ' {
				word = word + string(k)
			}
		}
	}
	str = append(str, word)

	return str
}
