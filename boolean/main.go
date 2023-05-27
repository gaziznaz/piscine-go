package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven() int {
	arg := os.Args[1:]
	if len(arg)%2 == 0 {
		return 1
	}
	return 0
}

func main() {
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if isEven() == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
