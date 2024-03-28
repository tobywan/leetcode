package pascalstriangle2

import "math/big"

func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)

	for i := 0; i <= rowIndex; i++ {
		ret[i] = nCr(rowIndex, i)
	}
	return ret
}

// {\displaystyle {n \choose k}={\frac {n!}{k!(n-k)!}}}

// nCr calculates n!/(r!(n-r)!)
// e.g. 33c10 = 33!/10!*23!)
// == 33*32*31*...*25*24/10*9*8*...*2*1
func nCr(n, r int) int {
	i := big.NewInt(0)
	i = i.Binomial(int64(n), int64(r))

	return int(i.Int64())
}
