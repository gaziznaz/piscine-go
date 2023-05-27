package piscine

func Capitalize(s string) string {
	srune := []rune(s)
	len := 0
	for j := range srune {
		len = j
	}
	for i := 0; i <= len; i++ {
		if srune[0] >= 97 && srune[0] <= 122 {
			srune[0] = srune[0] - 32
		}
		if i != len && ((srune[i] >= 0 && srune[i] <= 47) || (srune[i] >= 58 && srune[i] <= 64) ||
			(srune[i] >= 91 && srune[i] <= 96) || (srune[i] >= 123 && srune[i] <= 126)) {
			if srune[i+1] >= 97 && srune[i+1] <= 122 {
				srune[i+1] = srune[i+1] - 32
			}
		} else {
			if i != len && (srune[i+1] >= 65 && srune[i+1] <= 90) {
				srune[i+1] = srune[i+1] + 32
			}
		}
	}
	return string(srune)
}
