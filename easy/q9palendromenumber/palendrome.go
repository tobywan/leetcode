package q9panedromenumber

// isPalendrome1 builds an int iin reverse
func isPalindrome1(x int) bool {

	if x < 0 {
		return false
	}

	forwards := x
	reverse := 0

	for x != 0 {
		y := x / 10
		digit := x - y*10
		reverse = reverse*10 + digit
		x = y
	}

	return forwards == reverse
}

// isPalindrome2 uses a slice to compare
func isPalindrome2(x int) bool {

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
