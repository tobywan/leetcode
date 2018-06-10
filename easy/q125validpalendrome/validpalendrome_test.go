package validpalendrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{
			in:   "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			in:   "race a car",
			want: false,
		},
		{
			in:   "",
			want: true,
		},
	}

	for _, test := range tests {
		got := isPalindrome(test.in)
		if got != test.want {
			t.Errorf("isPalendrome(%q)=%t, want %t", test.in, got, test.want)
		}
	}

}

func TestFix(t *testing.T) {
	tests := []struct {
		in    rune
		want1 rune
		want2 bool
	}{
		{
			in:    '0',
			want1: '0',
			want2: true,
		},
		{
			in:    '5',
			want1: '5',
			want2: true,
		},
		{
			in:    '9',
			want1: '9',
			want2: true,
		},
		{
			in:    'a',
			want1: 'a',
			want2: true,
		},
		{
			in:    'm',
			want1: 'm',
			want2: true,
		},
		{
			in:    'z',
			want1: 'z',
			want2: true,
		},
		{
			in:    'A',
			want1: 'a',
			want2: true,
		},
		{
			in:    'B',
			want1: 'b',
			want2: true,
		},
		{
			in:    'Z',
			want1: 'z',
			want2: true,
		},
		{
			in:    '.',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    '@',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    ':',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    '[',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    '\u0060',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    '|',
			want1: '\u0000',
			want2: false,
		},
		{
			in:    '\u4845',
			want1: '\u0000',
			want2: false,
		},
	}
	for _, test := range tests {
		got1, got2 := fix(test.in)
		if got1 != test.want1 || got2 != test.want2 {
			t.Errorf("fix(%c)=%c,%t want %c,%t", test.in, got1, got2, test.want1, test.want2)
		}
	}
}
