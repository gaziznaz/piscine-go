package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	srune := []rune(s)
	for i := range srune {
		z01.PrintRune(srune[i])
	}
}
