package climbingstairs

func climbStairs(n int) int {
	a := 1
	b := 1

	for n > 0 {
		a, b = b, a+b
		n--
	}

	return a

}

func climbStairsJournal(n int) int {

	cache := make(intMap)

	return climbStair(n, cache)

}

type intMap map[int]int

func climbStair(n int, cache intMap) int {

	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	ret, ok := cache[n]

	if !ok {
		ret = climbStair(n-1, cache) + climbStair(n-2, cache)
		cache[n] = ret
	}
	// } else {
	// 	fmt.Printf("Cache hit for %d\n", n)
	// }

	return ret
}
