package timetosell

func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	buy := prices[0]

	for _, n := range prices[1:] {
		if n == buy {
			continue
		}
		if n > buy {
			profit = max(profit, n-buy)
		} else {
			buy = n
		}
	}
	return profit
}

func maxProfit1(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	buy := prices[0]

	for _, n := range prices[1:] {
		switch {

		case n > buy:
			profit = max(profit, n-buy)

		case n < buy:
			buy = n
		}
	}
	return profit
}

func max(m, n int) int {
	if m > n {
		return m
	}
	return n
}
