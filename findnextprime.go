package piscine

func IsPrime1(nb int) bool {
	if nb < 0 || nb == 0 || nb == 1 {
		return false
	} else {
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}

func FindNextPrime(nb int) int {
	if nb < 0 || nb == 0 || nb == 1 {
		return 2
	} else {
		for i := 0; i <= nb; i++ {
			if IsPrime1(nb) == false {
				nb++
			} else {
				break
			}
		}
	}
	return nb
}
