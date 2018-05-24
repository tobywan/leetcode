package twosum

func twoSum(nums []int, target int) []int {
	// read into map with index:
	idx := make(map[int][]int, len(nums))

	for i, n := range nums {
		idx[n] = append(idx[n], i)
	}

	for i, n := range nums {
		diff := target - n

		if ixs, ok := idx[diff]; ok {
			if len(ixs) > 1 {
				// we've been told only one answer, so
				// diff must == n
				return ixs[:2]
			}
			j := ixs[0]
			if i == j {
				continue //ignore same index
			}
			return []int{i, j}

		}

	}
	return []int{-1, -1}
}
