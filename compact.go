package piscine

func Compact(ptr *[]string) int {
	size := 0
	for _, i := range *ptr {
		if i != "" {
			size++
		}
	}
	arg := make([]string, size)
	k := 0
	for _, i := range *ptr {
		if i != "" {
			arg[k] = i
			k++
		}
	}
	*ptr = arg
	return size
}
