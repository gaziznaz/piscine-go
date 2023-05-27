package piscine

func RecursiveFactorial(nb int) int {
	a := 1
	if nb < 0 || nb > 25 {
		return 0
	} else if nb == 0 || nb == 1 {
		return 1
	} else {
		a = nb * RecursiveFactorial(nb-1)
		return a
	}
}
