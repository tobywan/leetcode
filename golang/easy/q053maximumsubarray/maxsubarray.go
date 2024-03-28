package maximumsubarray

// maxSubArray - Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	return maxSubArraySequential(nums)
}

// maxSubArray - Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and return its sum.
func maxSubArraySequential(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	msf := nums[0] // max so far
	meh := msf     // max ending here

	for _, n := range nums[1:] {
		meh = max(n, n+meh)
		msf = max(msf, meh)
	}

	return msf

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
