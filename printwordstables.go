package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := range a {
		for _, k := range a[i] {
			z01.PrintRune(k)
		}
		z01.PrintRune('\n')

	}
}
