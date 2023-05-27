package piscine

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s, toFind string) int {
	sub := []rune(toFind)
	str := []rune(s)
	lensub := 0
	lenstr := 0

	for i := range sub {
		lensub = i + 1
	}
	for j := range str {
		lenstr = j + 1
	}
	switch {
	case lensub == 0:
		return 0
	case lensub == 1:
		return IndexRune(str, sub[0])
	case lensub == lenstr:
		if s == toFind {
			return 0
		}
		return -1
	case lensub > lenstr:
		return -1
	case lensub < lenstr:
		c0 := sub[0]
		c1 := sub[1]
		i := 0
		t := lenstr - lensub + 1
		for i < t {
			if str[i] != c0 {
				o := IndexRune(str, c0)
				if o < 0 {
					return -1
				}
				i += o

			}
			if str[i+1] == c1 && s[i:i+lensub] == toFind {
				return i
			}
			i++

		}
	}
	return -1
}

func IndexRune(s []rune, r rune) int {
	index := -1
	for i := range s {
		if s[i] == r {
			index = i
			return index
		}
	}
	return index
}
