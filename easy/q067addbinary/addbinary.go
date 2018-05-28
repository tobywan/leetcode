package addbinary

func addBinary(a string, b string) string {

	la := len(a)
	lb := len(b)
	lmax := max(la, lb)

	result := make([]byte, lmax+1) // big enough to hold a flowed answer
	carry := off
	sum := off

	// iterate backwards over the input strings and the result byte slice concurrently
	for pos, posA, posB := lmax, la-1, lb-1; pos > 0; pos, posA, posB = pos-1, posA-1, posB-1 {

		bitA := off
		bitB := off
		if posA >= 0 {
			bitA = a[posA]
		}
		if posB >= 0 {
			bitB = b[posB]
		}
		sum, carry = addBits(bitA, bitB, carry)
		result[pos] = sum

	}

	result[0] = carry

	if carry == on {
		// We need the overflow bit too
		return string(result)
	}

	return string(result[1:])

}

// to facilitate conversion from bits to bytes
const (
	on  = byte('1')
	off = byte('0')
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// add three bits together and return 2 bits, sum is LSB and flow is MSB
func addBits(a, b byte, carry byte) (sum byte, flow byte) {
	const off3 = 3 * off
	sum = a + b + carry - off3

	if sum > 1 {
		sum = (sum & 1) + off
		flow = on
		return
	}
	flow = off
	sum = sum + off
	return
}
