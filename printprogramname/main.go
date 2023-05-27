package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for j, i := range arg[0] {
		if j > 1 {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
