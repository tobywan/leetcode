package lengthlastword

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{
			in:   "\U0001D400\U0001D401\U0001D402\U0001D403\U0001D404 \U0001D405\U0001D406\U0001D407\U0001D408\U0001D409",
			want: 5,
		},
		{
			in:   " xyz ",
			want: 3,
		},
		{
			in:   "xyz",
			want: 3,
		},
		{
			in:   " ",
			want: 0,
		},
		{
			in:   "a ",
			want: 1,
		},
		{
			in:   "",
			want: 0,
		},
		{
			in:   "hello world!",
			want: 6,
		},
	}

	for _, test := range tests {
		got := lengthOfLastWord(test.in)
		if got != test.want {
			t.Errorf("lengthOfLastWord(%q)=%d, want %d", test.in, got, test.want)
		}
	}
}
