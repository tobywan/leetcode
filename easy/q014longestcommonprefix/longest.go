package longestcommonprefix

// // longestCommonPrefix looks at successive initials
// // to see which are in common
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}
// 	if len(strs) == 1 {
// 		return strs[0]
// 	}
// 	var b bytes.Buffer
// 	pos := 0
// 	cont := true

// 	for cont && pos < len(strs[0]{
// 		r := strings. (strs[0], pos)
// 		for _, s := range strs {
// 			if len(s) >= pos || r != s[pos]{
// 				cont = false
// 				break
// 			}
// 		}
// 		if cont {
// 			b.WriteRune(r)
// 			pos ++
// 		}
// 	}

// 	return b.String()
// }

// longestCommonPrefix stores a prefix as runes and compares
// each string against them

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	pref := []rune(strs[0])

	for _, s := range strs[1:] {
		pref = commonPrefix(pref, s)
	}

	return string(pref)
}

// commonPrefix finds the prefix runes p has in common with s
func commonPrefix(p []rune, s string) []rune {
	l := len(p)
	if l == 0 {
		return p
	}

	common := 0
	for n, r := range s {
		if n >= l || p[n] != r {
			break
		}
		common++
	}
	return p[:common]

}

// func longestCommonPrefix(strs []string) string {

// 	l := &lexTree{}

// 	for _, s := range strs {
// 		l.append([]rune(s))
// 	}
// 	return l.commonPrefix()

// }

// type lexTree struct {
// 	kids map[rune](*lexTree)
// }

// // append looks to see if s can be added to any kids
// func (l *lexTree) append(s []rune) {
// 	if len(s) == 0 {
// 		return
// 	}

// 	if l.kids == nil {
// 		l.kids = make(map[rune]*lexTree)
// 	}
// 	var lt lexTree
// 	if lt, ok := l.kids[s[0]]; !ok {
// 		// already have an entry for initial letter
// 		lt = &lexTree{}
// 		l.kids[s[0]] = lt
// 	}
// 	lt.append(s[1:])
// }

// func (l *lexTree) commonPrefix() string {
// 	var b bytes.Buffer

// 	x := len(l.kids)
// 	kid := l
// 	for x == 1 {
// 		for k, v := range kid.kids {
// 			b.WriteRune(k)
// 			kid = v
// 		}
// 		x = len(kid.kids)

// 	}

// 	return b.String()
// }
