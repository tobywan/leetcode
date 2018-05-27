package implementstrstr

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		hs   string
		nd   string
		want int
	}{
		{
			hs:   "a",
			nd:   "aaaaaaaaa",
			want: -1,
		},
		{
			hs:   "aaaaaaaaaab",
			nd:   "b",
			want: 10,
		},
		{
			hs:   "abcabcabcabcXabcabcabcX",
			nd:   "abcX",
			want: 9,
		},
		{
			hs:   "ababababababXababab",
			nd:   "ababX",
			want: 8,
		},
		{
			hs:   "aaaaa",
			nd:   "bba",
			want: -1,
		},
		{
			hs:   "\u1230\u1231\u1232\u1233\u1234\u1235\u1236",
			nd:   "\u1233\u1234",
			want: 3,
		},
		{
			hs:   "\u1230\u1231\u1232\u1233\u1234\u1235\u1236",
			nd:   "\u0034",
			want: -1,
		},
		{
			hs:   "hello",
			nd:   "ll",
			want: 2,
		},
		{
			hs:   "hello",
			nd:   "",
			want: 0,
		},
	}

	for _, test := range tests {
		got := strStr(test.hs, test.nd)
		if got != test.want {
			t.Errorf("strStr(%q,%q)=%d, want %d", test.hs, test.nd, got, test.want)
		}
	}
}
