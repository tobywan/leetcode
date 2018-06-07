package timetosell2

func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}

	totalProfit := 0
	subProfit := 0
	buy := prices[0]

	for i, this := range prices[1:] {
		last := prices[i]
		if last == this {
			// no change in price
			continue
		}

		if last < this {
			// a profit to be made
			subProfit = this - buy
		} else {
			// prices have gone down since yesterday
			buy = this
			totalProfit += subProfit
			subProfit = 0
		}
	}
	totalProfit += subProfit

	return totalProfit

}
