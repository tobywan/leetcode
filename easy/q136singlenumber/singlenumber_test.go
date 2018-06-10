package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {

	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			in:   []int{2, 2, 1},
			want: 1,
		},
		{
			in:   []int{2, 2, 1, 3, 3},
			want: 1,
		},
	}
	for _, test := range tests {
		got := singleNumber(test.in)
		if got != test.want {
			t.Errorf("singleNumber(%v)=%d, want %d", test.in, got, test.want)
		}
	}

}
