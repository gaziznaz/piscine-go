package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := range arg {
		if i > 0 {
			for _, k := range arg[i] {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
