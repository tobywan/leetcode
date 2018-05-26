package removeduplicates

func removeDuplicates(nums []int) int {

	l := len(nums)

	if len(nums) <= 1 {
		return l
	}

	p1, p2 := 0, 1

	for p2 < l {
		if nums[p1] != nums[p2] {
			nums[p1+1] = nums[p2]
			p1++
		}
		p2++
	}

	return p1 + 1
}
