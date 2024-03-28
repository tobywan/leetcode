package removeduplicates

func removeDuplicates(nums []int) int {

	l := len(nums)

	if len(nums) <= 1 {
		return l
	}

	p1 := 0

	for p2 := 1; p2 < l; p2++ {
		if nums[p1] != nums[p2] {
			nums[p1+1] = nums[p2]
			p1++
		}
	}

	return p1 + 1
}
