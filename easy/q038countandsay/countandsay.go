package countandsay

import (
	"bytes"
	"strconv"
)

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	s := sentence{
		num: 2,
		sequences: []sequence{
			{1, 1},
		},
	}
	for s.num < n {
		s.increment()
	}
	return s.String()
}

// sentence translates the number to its phrase
type sentence struct {
	num       int
	sequences []sequence
}

func (s *sentence) String() string {
	var b bytes.Buffer
	for _, seq := range s.sequences {

		b.WriteString(seq.String())
	}
	return b.String()
}

func (s *sentence) increment() {
	say := s.String()
	s.num++
	s.sequences = []sequence{}
	r1 := rune(say[0])
	count := 1
	for _, r2 := range say[1:] {
		if r1 == r2 {
			count++
		} else {
			// flush the previous rune
			s.sequences = append(s.sequences, sequence{int(r1 - '0'), count})
			r1 = r2
			count = 1
		}
	}
	// flush the last run
	s.sequences = append(s.sequences, sequence{int(r1 - '0'), count})

}

type sequence struct {
	digit int
	count int
}

func (s *sequence) String() string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(s.count))
	b.WriteString(strconv.Itoa(s.digit))

	return b.String()
}
