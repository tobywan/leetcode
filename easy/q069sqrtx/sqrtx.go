package sqrtx

func mySqrt(x int) int {

	if x <= 1 {
		return x
	}
	b := bounded(x)

	return narrow(b, x)
}

func bounded(x int) *bounds {

	lb := 1
	ub := x

	lower := lb*lb <= x
	for lower {
		t := lb * 2
		sqr := t * t

		lower = sqr <= x
		if lower {
			lb = t
		} else {
			ub = t
		}

	}
	return &bounds{lb, ub}

}

type bounds struct {
	low  int
	high int
}

func (b *bounds) mid() int {
	return (b.high + b.low) / 2
}

// narrow the bounds to find the sqrt
func narrow(b *bounds, x int) int {

	m := b.mid()
	for (m > b.low) && (b.low*b.low <= x) && (b.high*b.high > x) {
		if m*m > x {
			b.high = m
		} else {
			b.low = m
		}
		m = b.mid()
	}

	return b.low

}

func mySqrt2(x int) int {

	if x <= 1 {
		return x
	}

	lb, found := lowerbound(x)
	if found {
		return lb
	}

	b := &bounds{lb, x}
	return narrow(b, x)

}

// find a course lower bound for sqrt.
// If we accidentally find it in the process we'll return true
func lowerbound(x int) (lb int, foundit bool) {

	lb1 := 1

	lower := lb1*lb1 < x
	for lower {
		lb2 := lb1 * 2
		sqr := lb2 * lb2
		if sqr == x {
			return lb2, true
		}
		lower = sqr < x
		if lower {
			lb1 = lb2
		}

	}
	return lb1, false
}
