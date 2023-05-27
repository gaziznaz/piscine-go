package piscine

func Max(arr []int) int {
	max := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
