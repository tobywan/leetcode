package majorityelement

func majorityElement(nums []int) int {

	counts := make(map[int]int)
	maxCount := 0
	major := 0

	for _, n := range nums {
		count := counts[n] + 1
		if count > maxCount {
			maxCount = count
			major = n
		}
		counts[n] = count
	}

	return major
}
