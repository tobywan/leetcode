package q9panedromenumber

import (
	"testing"
)

func TestIsPalendrome(t *testing.T) {
	//	t.Logf(" 2^31 = %d", (1 << 31))
	tests := []struct {
		in   int
		want bool
	}{
		{
			in:   1,
			want: true,
		},
		{
			in:   -1,
			want: false,
		},
		{
			in:   0,
			want: true,
		},
		{
			in:   1234567,
			want: false,
		},
		{
			in:   12345654321,
			want: true,
		},
		{
			in:   11111111111,
			want: true,
		},
		{
			in:   99,
			want: true,
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.in)
		if got != test.want {
			t.Errorf("isPalindrome1(%d)=%t, want %t", test.in, got, test.want)
		}
	}
}
