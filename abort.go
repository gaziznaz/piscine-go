package piscine

func Abort(a, b, c, d, e int) int {
	arr := [5]int{a, b, c, d, e}
	var temp int
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr[2]
}
