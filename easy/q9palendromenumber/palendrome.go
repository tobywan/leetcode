package q9panedromenumber

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	digits := []int{}
	for x != 0 {
		y := x / 10
		digit := x - y*10
		digits = append(digits, digit)
		x = y
	}

	l := len(digits)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
