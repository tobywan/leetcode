package removeelement

func removeElement(nums []int, val int) int {

	l := len(nums)

	if l == 0 {
		return 0
	}

	ret := nums[:0]
	for _, n := range nums {
		if n != val {
			ret = append(ret, n)
		}
	}

	return len(ret)
}
