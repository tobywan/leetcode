package longestcommonprefix

import (
	"testing"
)

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
