package countandsay

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{
			in:   1,
			want: "1",
		},
		{
			in:   2,
			want: "11",
		},
		{
			in:   3,
			want: "21",
		},
		{
			in:   4,
			want: "1211",
		},
		{
			in:   5,
			want: "111221",
		},
		{
			in:   6,
			want: "312211",
		},
		{
			in:   7,
			want: "13112221",
		},
		{
			in:   8,
			want: "1113213211",
		},
		{
			in:   9,
			want: "31131211131221",
		},
		{
			in:   10,
			want: "13211311123113112211",
		},
	}
	for _, test := range tests {
		got := countAndSay(test.in)
		if got != test.want {
			t.Errorf("countAndSay(%d)=%q, want %q", test.in, got, test.want)
		}
	}
}
