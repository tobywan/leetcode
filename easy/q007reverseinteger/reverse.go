package reverseinteger

func reverse(x int) int {

	const minInt = -(1 << 31)
	const maxInt = (1 << 31) - 1
	if -10 < x && x < 10 {
		return x
	}

	digits := []int{}
	for x != 0 {
		y := x / 10
		digit := x - y*10
		digits = append(digits, digit)
		x = y
	}
	ret := 0
	for _, digit := range digits {
		ret = ret*10 + digit
	}

	if ret < minInt || ret > maxInt {
		return 0
	}
	return ret
}
