package main

import "github.com/01-edu/z01"

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func PrintMemory(arr [10]int) {
	var res []string
	for i := 0; i < 10; i++ {
		el := ToHex(arr[i])
		for i := len(ToHex(arr[i])); i < 4; i++ {
			el += "0"
		}
		res = append(res, el)
		res = append(res, "0000")
	}
	PrintRes(res)
	for k := range arr {
		z01.PrintRune(rune(arr[k]))
		if rune(arr[k]) < 31 {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func ToHex(n int) string {
	base := "0123456789abcdef"
	if n < 16 {
		return base[n : n+1]
	}
	return ToHex(n/16) + base[n%16:n%16+1]
}

func PrintRes(s []string) {
	for i := range s {
		for _, k := range s[i] {
			z01.PrintRune(rune(k))
		}
		if i != 0 && (i+1)%8 == 0 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
