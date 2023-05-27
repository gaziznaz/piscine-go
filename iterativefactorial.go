package piscine

func IterativeFactorial(nb int) int {
	a := 1
	if (nb < 0) || (nb > 50) {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := 1; i <= nb; i++ {
			a = i * a
		}
	}
	if a < 0 {
		return 0
	}
	return a
}
