package piscine

func CollatzCountdown(start int) int {
	res := 0
	if start <= 0 {
		res = -1
	} else {
		for {
			if start != 1 {
				if start%2 == 0 {
					start = start / 2
				} else {
					start = start*3 + 1
				}
				res++
			} else {
				break
			}
		}
	}
	return res
}
