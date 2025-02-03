func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var left int = 0
	var right int = 1
	var profit_max int = 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[left] < prices[right] {
			var profit int = prices[right] - prices[left]
			if profit > profit_max {
				profit_max = profit
			}
		} else {
			left = right
		}
		right += 1
	}

	return profit_max
}
