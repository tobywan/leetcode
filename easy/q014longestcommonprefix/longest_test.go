package longestcommonprefix

import (
	"testing"
	"unicode/utf8"
)

func TestStringSlice(t *testing.T) {
	s := "\u1234\u2222"

	for i := 0; i <= len(s); i++ {
		tmp := s[:i]
		t.Logf("s[:%d] = %q, is Valid %t", i, tmp, utf8.ValidString(tmp))
	}

}
func BenchmarkLongestCommonPrefix(b *testing.B) {

	for i := 0; i < b.N; i++ {
		longestCommonPrefix([]string{"dogging", "doggy", "dogget"})
	}
}
func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		in   []string
		want string
	}{
		{
			in:   []string{"\u1234flower", "\u1234flow", "\u1234flight"},
			want: "\u1234fl",
		},
		{
			in:   []string{"fl\u2222ower", "fl\u2222ow", "fl\u2222ight"},
			want: "fl\u2222",
		},
		{
			in:   []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			in:   []string{"car"},
			want: "car",
		},
		{
			in:   []string{"1234567890", ""},
			want: "",
		},
		{
			in:   []string{"1234567890", "1234567890"},
			want: "1234567890",
		},
		{
			in:   []string{"dogging", "doggy", "dogget"},
			want: "dogg",
		},
		{
			in:   []string{"dog", "racecar", "car"},
			want: "",
		},
	}
	for _, test := range tests {
		got := longestCommonPrefix(test.in)
		if got != test.want {
			t.Errorf("longestCommonPrefix(%q)=%q, want %q", test.in, got, test.want)
		}
	}
}
