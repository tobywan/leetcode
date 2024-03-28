package factorialzeros

// number of trailing zeros corresponds to number of 5s
// result = (n % 5)
//  + (n % 25)
// + (n % 125)
//  + ...
func trailingZeroes(n int) int {
	result := 0
	divisor := 5
	quot := n / divisor

	for quot > 0 {
		result += quot
		divisor *= 5
		quot = n / divisor
	}

	return result
}
