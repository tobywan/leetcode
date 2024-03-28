package twosum

func twoSum(numbers []int, target int) []int {

	seen := make(map[int]int, len(numbers))

	// we can fast forward to the 1st candidate for a sum.
	// The lower value is

	for i, n := range numbers {
		seek := target - n
		j, ok := seen[seek]
		if ok {
			return []int{j + 1, i + 1}
		}
		seen[n] = i
	}

	return []int{}
}
