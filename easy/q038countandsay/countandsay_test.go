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
		{
			in:   20,
			want: "11131221131211132221232112111312111213111213211231132132211211131221131211221321123113213221123113112221131112311332211211131221131211132211121312211231131112311211232221121321132132211331121321231231121113112221121321133112132112312321123113112221121113122113121113123112112322111213211322211312113211",
		},
	}
	for _, test := range tests {
		got := countAndSay(test.in)
		if got != test.want {
			t.Errorf("countAndSay(%d)=%q, want %q", test.in, got, test.want)
		}
	}
}
