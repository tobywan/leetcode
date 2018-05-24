package q7reverseinteger

import (
	"testing"
)

func TestAdhoc(t *testing.T) {
	t.Logf("%d", reverse(1234567))
}

func TestReverse(t *testing.T) {
	//	t.Logf(" 2^31 = %d", (1 << 31))
	tests := []struct {
		in   int32
		want int32
	}{
		{
			in:   1,
			want: 1,
		},
		{
			in:   -1,
			want: -1,
		},
		{
			in:   0,
			want: 0,
		},
		{
			in:   1234567,
			want: 7654321,
		},
		{
			in:   -1000001,
			want: -1000001,
		},
		{
			in:   -1111199999,
			want: -0,
		},
		{
			in:   1111199999,
			want: 0,
		},
	}

	for _, test := range tests {
		got := reverse(int(test.in))
		if got != int(test.want) {
			t.Errorf("reverse(%d)=%d, want %d", test.in, got, test.want)
		}
	}
}
