package validpalendrome

func isPalindrome(s string) bool {

	if s == "" {
		return true
	}
	// build a lowercase/numeric slice from the string
	// Over allocate storage if needed to save re-allocation later
	ln := make([]rune, 0, len(s))

	for _, r := range s {

		rok, ok := fix(r)
		if !ok {
			continue
		}
		ln = append(ln, rok)
	}

	// Now check if it's a palendrome
	for i, j := 0, len(ln)-1; i < j; i, j = i+1, j-1 {
		if ln[i] != ln[j] {
			return false
		}
	}

	return true
}

// Fix checks to see if the character is alpha numeric and
// lower cases at the same time.
func fix(r rune) (okrune rune, ok bool) {
	if ('0' <= r && r <= '9') || ('a' <= r && r <= 'z') {
		return r, true
	}

	if 'A' <= r && r <= 'Z' {
		return r - 'A' + 'a', true
	}

	return '\u0000', false
}
