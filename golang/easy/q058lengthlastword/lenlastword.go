package lengthlastword

func lengthOfLastWord(s string) int {

	sr := []rune(s)

	right := 0
	left := 0
	inword := false

	for pos := len(sr) - 1; pos >= 0; pos-- {
		if sr[pos] == ' ' {
			if inword {
				left = pos
				inword = false
				break
			}
		} else {
			if !inword {
				inword = true
				right = pos
			}
		}
	}

	if inword {
		// We hit the beginning without a closing space
		left--
	}

	return right - left
}
