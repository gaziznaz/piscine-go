package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	len := 0
	for i := range arg {
		len = i
	}
	for j := len; j >= 0; j-- {
		if j > 0 {
			for _, k := range arg[j] {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
