package implementstrstr

import "unicode/utf8"

func strStr(haystack string, needle string) int {

	ln := len(needle)
	lh := len(haystack)

	for pos := 0; pos+ln <= lh; pos++ {

		if begins(haystack[pos:], needle) {
			return utf8.RuneCountInString(haystack[:pos])
		}

	}
	return -1
}

func begins(s, prefix string) bool {

	if len(prefix) > len(s) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}
