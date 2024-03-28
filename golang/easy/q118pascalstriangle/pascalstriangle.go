package pascalstriangle

func generate(numRows int) [][]int {

	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			res[0] = []int{1}
			continue
		}
		res[i] = nextRow(res[i-1])

	}

	return res
}

func nextRow(row []int) []int {

	l := len(row)

	if l == 1 {
		return []int{1, 1}
	}

	ret := make([]int, l+1)
	ret[0] = 1
	ret[l] = 1

	mid := (l + 1) / 2

	for i := 1; i <= mid; i++ {
		ret[i] = row[i-1] + row[i]
		ret[l-i] = ret[i]
	}

	return ret

}
