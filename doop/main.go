package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) != 4 {
	} else if !(CheckValue(arg[1])) {
		return
	} else if !(CheckValue(arg[3])) {
		return
	} else {
		res := 0
		a := ToInt(arg[1])
		b := ToInt(arg[3])
		if !(CheckOverflow(a)) {
			return
		}
		if !(CheckOverflow(b)) {
			return
		}

		if arg[2] == "+" {
			res = a + b
		} else if arg[2] == "*" {
			res = a * b
		} else if arg[2] == "-" {
			res = a - b
		} else if arg[2] == "%" {
			if b == 0 {
				PrintErrMod()
				return
			} else {
				res = a % b
			}
		} else if arg[2] == "/" {
			if b == 0 {
				PrintErrDiv()
				return
			} else if (a < b) && a*(-1) > b {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return

			} else {
				res = a / b
			}
		} else {
			return
		}
		if !(CheckOverflow(res)) {
			return
		}
		result := ToRune(res)
		if res < 0 {
			result = append(result, '-')
		}
		for i := len(result) - 1; i >= 0; i-- {
			z01.PrintRune(result[i])
		}
		z01.PrintRune('\n')
	}
}

func PrintErrMod() {
	s := "No modulo by 0"
	for _, k := range s {
		z01.PrintRune(rune(k))
	}
	z01.PrintRune('\n')
}

func PrintErrDiv() {
	s := "No division by 0"
	for _, k := range s {
		z01.PrintRune(rune(k))
	}
	z01.PrintRune('\n')
}

func CheckValue(s string) bool {
	srune := []rune(s)
	for _, k := range srune {
		if k != 45 && (k < 48 || k > 59) {
			return false
		}
	}
	return true
}

func ToInt(s string) int {
	var num int
	i := 1
	for _, k := range s {
		if k >= 48 && k <= 57 {
			num = num*10 + (int(k) - 48)
		} else if k == 45 {
			i = -1
		}
	}
	return num * i
}

func CheckOverflow(num int) bool {
	if num < (-9223372036854775808) || num > 9223372036854775807 {
		return false
	}
	return true
}

func ToRune(num int) []rune {
	var result []rune
	for {
		if num < 10 && num > 0 {
			result = append(result, rune(num+48))
			break
		} else if num < 0 {
			num = num * (-1)
		} else {
			result = append(result, rune(num%10+48))
			num = num / 10
		}
	}
	return result
}
