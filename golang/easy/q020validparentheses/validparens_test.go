package validparentheses

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "",
			want: true,
		},
		{
			in:   "()",
			want: true,
		},
		{
			in:   "()[]{}",
			want: true,
		},
		{
			in:   "(",
			want: false,
		},
		{
			in:   "]",
			want: false,
		},
		{
			in:   "(]",
			want: false,
		},
		{
			in:   "([)]",
			want: false,
		},
		{
			in:   "{[]}",
			want: true,
		},
	}
	for _, test := range tests {
		got := isValid(test.in)
		if got != test.want {
			t.Errorf("isValid{%q) = %t, want %t", test.in, got, test.want)
		}
	}
}
