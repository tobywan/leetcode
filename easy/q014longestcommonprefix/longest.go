package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	pref := strs[0][:]

	for _, s := range strs[1:] {
		pref = commonPrefix(pref, s)
	}

	return pref
}

// commonPrefix finds the prefix p has in common with s
func commonPrefix(p string, s string) string {
	lp := len(p)
	ls := len(s)
	if lp == 0 || ls == 0 {
		return ""
	}

	var lc int // length of shorter string
	if lp < ls {
		lc = lp
	} else {
		lc = ls
	}

	var cb int // common bytes

	for cb = 0; cb < lc; cb++ {
		if p[cb] != s[cb] {
			break
		}
	}

	return p[:cb]

}
