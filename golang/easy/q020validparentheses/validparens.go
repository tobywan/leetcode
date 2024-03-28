package validparentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]rune, 0, len(s))

	for _, r := range s {
		if _, ok := closers[r]; !ok {
			stack = push(stack, r)
		} else {
			stack, ok = pop(stack, r)
			if !ok {
				return false
			}
		}
	}

	return len(stack) == 0
}

var closers = map[rune]rune{
	'}': '{',
	']': '[',
	')': '(',
}

// push adds an opener to the stack
func push(stack []rune, r rune) []rune {
	return append(stack, r)
}

// pop closes the opener or returns false if that's invalid
func pop(stackin []rune, r rune) (stackout []rune, ok bool) {
	l := len(stackin)
	if l == 0 {
		return stackin, false
	}
	c := closers[r]
	if c != stackin[l-1] {
		return stackin, false
	}
	return stackin[:l-1], true
}
