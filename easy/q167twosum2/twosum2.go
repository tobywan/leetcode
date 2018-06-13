package twosum

func twoSum(numbers []int, target int) []int {

	seen := make(map[int]int)

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
