package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 && IsValidValue(args[0]) && IsValidOperator(args[1]) && IsValidValue(args[2]) {
		Calculate(Atoi, args[0], args[1], args[2])
	}
}

func PrintStr(s string) {
	sByByte := []byte(s)
	for _, v := range sByByte {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	var arrayOfInts []int

	if n == 0 {
		arrayOfInts = append(arrayOfInts, 0)
	} else if n < 0 {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, -(n % 10))
			n /= 10
		}
		z01.PrintRune('-')
	} else {
		for n != 0 {
			arrayOfInts = append(arrayOfInts, n%10)
			n /= 10
		}
	}

	for i := len(arrayOfInts) - 1; i >= 0; i-- {
		z01.PrintRune(rune(arrayOfInts[i] + 48))
	}
	z01.PrintRune('\n')
}

func IsValidValue(s string) bool {
	for _, v := range s {
		if (v < '0' || v > '9') && v != '-' {
			return false
		}
	}
	return true
}

func IsValidOperator(s string) bool {
	for _, v := range s {
		if v == '+' || v == '-' || v == '/' || v == '*' || v == '%' {
			return true
		}
	}
	return false
}

func Calculate(f func(s string) (int, bool), first, operator, second string) {
	a, aBool := f(first)
	b, bBool := f(second)
	if aBool && bBool {
		switch operator {
		case "+":
			c := a + b
			if (c > a) == (b > 0) {
				PrintNbr(c)
			}
		case "-":
			c := a - b
			if (c < a) == (b > 0) {
				PrintNbr(c)
			}
		case "*":
			c := a * b
			if (c < 0) == ((a < 0) != (b < 0)) {
				if c/b == a {
					PrintNbr(c)
				}
			}
		case "/":
			if b != 0 {
				c := a / b
				PrintNbr(c)
			} else {
				PrintStr("No division by 0")
			}
		case "%":
			if b != 0 {
				c := a % b
				PrintNbr(c)
			} else {
				PrintStr("No modulo by 0")
			}
		}
	}
}

func Atoi(s string) (int, bool) {
	sum := 0
	if len(s) > 0 {
		for i, v := range s {
			if i == 0 && v == '+' {
			} else if i == 0 && v == '-' {
			} else if v >= '0' && v <= '9' {
				c := sum*10 + int(v-48)
				if (c > sum*10) == (int(v-48) > 0) {
					sum = c
				} else {
					return 0, false
				}
			} else {
				return 0, false
			}
		}
		if s[0] == '-' {
			return -sum, true
		} else {
			return sum, true
		}
	}
	return 0, false
}
