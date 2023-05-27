package piscine

func SortIntegerTable(table []int) {
	tmp := 0
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-1; j++ {
			if table[j] > table[j+1] {
				tmp = table[j]
				table[j] = table[j+1]
				table[j+1] = tmp
			}
		}
	}
}
