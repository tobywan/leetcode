package excelnumber

import "math"

func titleToNumber(s string) int {
	l := len(s)
	digit := 0
	result := 0

	for i := l - 1; i >= 0; i-- {
		b := s[i]
		placeVal := math.Pow(26.0, float64(digit))
		result += val(b) * int(placeVal)
		digit++
	}

	return result
}

func val(b byte) int {
	return int(b + 1 - byte('A'))
}
