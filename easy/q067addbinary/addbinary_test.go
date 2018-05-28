package addbinary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {

	tests := []struct {
		in1  string
		in2  string
		want string
	}{
		{
			in1:  "100000",
			in2:  "11111",
			want: "111111",
		},
		{
			in1:  "111111",
			in2:  "1",
			want: "1000000",
		},
		{
			in1:  "111",
			in2:  "111",
			want: "1110",
		},
		{
			in1:  "101010101",
			in2:  "10101010",
			want: "111111111",
		},
		{
			in1:  "11",
			in2:  "1",
			want: "100",
		},
	}

	for _, test := range tests {
		got := addBinary(test.in1, test.in2)
		if got != test.want {
			t.Errorf("addBinary(%q,%q)=%q, want %q", test.in1, test.in2, got, test.want)
		}
	}
}
