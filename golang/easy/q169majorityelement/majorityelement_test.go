package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			in:   []int{3, 2, 3},
			want: 3,
		},
		{
			in:   []int{1, 1, 1, 1, 1, 4, 4, 4, 4, 4, 4},
			want: 4,
		},
		{
			in:   []int{1, 1, 1, 1, 1, 4, 4, 4, 4},
			want: 1,
		},
	}
	for _, test := range tests {
		got := majorityElement(test.in)
		if test.want != got {
			t.Errorf("majorityElement(%v)=%d, want %d", test.in, got, test.want)
		}
	}

}
