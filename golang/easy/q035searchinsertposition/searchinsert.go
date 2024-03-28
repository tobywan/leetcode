package searchinsertposition

import "fmt"

// searchInsert more elegant solutin not using recursion
func searchInsert(nums []int, target int) int {
	size := len(nums)

	if size == 0 {
		return 0
	}

	right := size - 1
	left := 0

	// move start to the target or
	for left < right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}
	if nums[left] >= target {
		return left
	}

	return left + 1

}

func searchInsertRec(nums []int, target int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	if target <= nums[0] {
		return 0
	}

	if target > nums[l-1] {
		return l
	}
	// target is in nums or should be in between the values

	return findPos(nums, target, 0)

}

// findPos locates the position in the nums slice.
// offset is the index of the first element relative to the beginning of
// the original slice
// 2 4 6 8 10 12 14 16 18 20
func findPos(nums []int, target int, index int) int {

	//	fmt.Printf("nums: %v, index: %d\n", nums, index)
	i := len(nums) - 1
	if i < 1 {
		panic(fmt.Sprintf("Slice too short, must be at least 2:%v", nums))
	}
	// divide nums into 2

	j := i / 2

	// Check either side of the division

	if target == nums[j] {
		return index + j
	}
	if target == nums[j+1] {
		return index + j + 1
	}

	if target > nums[j] && target < nums[j+1] {
		return index + j + 1
	}

	if target < nums[j] {
		return findPos(nums[:j+1], target, index)
	}

	// only option, target is in greater half
	return findPos(nums[j:], target, index+j)

}
