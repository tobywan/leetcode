package plusone

func plusOne(digits []int) []int {

	l := len(digits)
	overflow := true
	for pos := l - 1; pos >= 0 && overflow; pos-- {
		digits[pos], overflow = inc(digits[pos])

	}

	if !overflow {
		return digits
	}

	ret := make([]int, 0, l+1)
	ret = append(ret, 1)
	ret = append(ret, digits...)
	return ret
}

func inc(digit int) (ans int, overflow bool) {
	if digit == 9 {
		return 0, true
	}
	return digit + 1, false
}
