package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	len := 0
	var temp string
	for i := range arg {
		len = i
	}
	for j := 0; j < len; j++ {
		for k := 1; k < len; k++ {
			if arg[k] > arg[k+1] {
				temp = arg[k]
				arg[k] = arg[k+1]
				arg[k+1] = temp
			}
		}
	}
	for l := range arg {
		if l > 0 {
			for _, k := range arg[l] {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
