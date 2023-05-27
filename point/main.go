package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func toRune(p int) []rune {
	var a []rune
	m := 1
	for i := 0; i < m; i++ {
		if p > 9 {
			a = append(a, rune(p/10)+48)
			p = p % 10
			m++
		} else {
			a = append(a, rune(p)+48)
		}
	}
	return a
}

func main() {
	points := point{}
	setPoint(&points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, i := range toRune(points.x) {
		z01.PrintRune(i)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	for _, i := range toRune(points.y) {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
