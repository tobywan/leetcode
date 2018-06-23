package exceltitle

import (
	"bytes"
	"math"
)

/*
 1   A           0*26 + 1
26   Z           0*26 + 26
27   AA          1*26 + 1
52   AZ          1*26 + 26
53   BA          2*26 + 1
701  ZY          26*26 + 25
702  ZZ          26*26 + 26
703 AAA  1*678 + 1*26 + 1
*/
var placeValues = make(map[int]int)

//  excelDigits finds the number of excel digits in a number
func excelDigits(n int) int {

	if n < 1 {
		return 0
	}
	aa := 1 // aa represents numbers that can be written with A's only
	digits := 0
	placeValue := 1

	for aa <= n {
		placeValues[digits] = placeValue
		digits++
		placeValue = int(math.Pow(26.0, float64(digits)))
		aa += placeValue

	}

	return digits
}

func convertToTitle(n int) string {
	digits := excelDigits(n)
	var b bytes.Buffer

	for i := digits - 1; i >= 0; i-- {
		pv := placeValues[i]
		val := n / pv
		n -= val * pv

		if i > 0 && n == 0 {
			// Special case. as we want Z=26
			val--
			n = 26
		}
		r := 'A' + rune(val-1)
		b.WriteRune(r)
	}
	return b.String()
}

// func convertToTitle(n int) string {
// 	reverse := make([]rune, 0, 10) // assume 10 letters is OK for now

// 	var r rune
// 	for n > 0 {
// 		r, n = units(n)
// 		reverse = append(reverse, r)
// 	}

// 	var b bytes.Buffer

// 	for i := len(reverse) - 1; i >= 0; i-- {
// 		b.WriteRune(reverse[i])
// 	}

// 	return b.String()
// }

// // units finds the letter represented by the "units" of the number in base 26
// // and returns its letter and also n >> 1 in base 26
// func units(n int) (letter rune, shifted int) {

// 	if n == 26 {
// 		return 'Z', 0
// 	}

// 	u := n % 26
// 	shift := n / 26

// 	return rune(u-1) + 'A', shift

// }
