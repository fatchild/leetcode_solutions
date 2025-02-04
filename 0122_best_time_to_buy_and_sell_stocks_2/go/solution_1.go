func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var profit_max int = 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			profit_max += prices[i+1] - prices[i]
		}
	}

	return profit_max
}
