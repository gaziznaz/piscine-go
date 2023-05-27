package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		num, check := Atoi(os.Args[1])

		if !check {
			return
		}
		if IsPrime(num) {
			PrintRes(num)
			z01.PrintRune('\n')
			return
		}
		var result []int
		for i := 2; i <= num; i++ {
			if IsPrime(i) && num%i == 0 {

				result = append(result, i)
				num = num / i
				i--

			}
		}
		for i, k := range result {
			PrintRes(k)
			if i != len(result)-1 {
				z01.PrintRune('*')
			}
		}
		if len(result) != 0 {
			z01.PrintRune('\n')
		}

	}
}

func Atoi(s string) (int, bool) {
	num := 0
	check := true

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(rune(s[i])-48)
		} else {
			check = false
		}
	}
	return num, check
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func PrintRes(n int) {
	var res []rune

	for n > 0 {
		res = append(res, rune(n%10)+48)
		n = n / 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		z01.PrintRune(res[i])
	}
}
