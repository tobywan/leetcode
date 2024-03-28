package sqrtx

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{
			in:   1000001,
			want: 1000,
		},
		{
			in:   8,
			want: 2,
		},
		{
			in:   144,
			want: 12,
		},
		{
			in:   0,
			want: 0,
		},
		{
			in:   4,
			want: 2,
		},
	}
	for _, test := range tests {
		got := mySqrt(test.in)
		if test.want != got {
			t.Errorf("sqrtx(%d)=%d, want %d", test.in, got, test.want)
		}
	}
}
