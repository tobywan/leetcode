package exceltitle

import "testing"

func TestExcelDigits(t *testing.T) {
	tests := []struct {
		in   int
		want int
	}{
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 1,
		},
		{
			in:   26,
			want: 1,
		},
		{
			in:   27,
			want: 2,
		},
		{
			in:   701,
			want: 2,
		},
		{
			in:   702,
			want: 2,
		},
		{
			in:   703,
			want: 3,
		},
		{
			in:   704,
			want: 3,
		},
	}
	for _, test := range tests {
		got := excelDigits(test.in)
		if test.want != got {
			t.Errorf("excelDigits(%d)=%d, want %d", test.in, got, test.want)
		}
	}

}

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		in   int
		want string
	}{
		{
			in:   1,
			want: "A",
		},
		{
			in:   28,
			want: "AB",
		},
		{
			in:   52,
			want: "AZ",
		},
		{
			in:   701,
			want: "ZY",
		},
		{
			in:   27,
			want: "AA",
		},
		{
			in:   703,
			want: "AAA",
		},
		{
			in:   -1,
			want: "",
		},
		{
			in:   702,
			want: "ZZ",
		},
	}

	for _, test := range tests {
		got := convertToTitle(test.in)
		if got != test.want {
			t.Errorf("convertToTitle(%d)=%q, want %q", test.in, got, test.want)
		}
	}

}
