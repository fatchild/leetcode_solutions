func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var buy int = prices[0]
	var sell int = prices[1]
	var buy_idx int = 0
	var sell_idx int = 1
	var low int = prices[0]
	var low_idx int = 0
	var profit int = sell - buy

	for i := 1; i < len(prices)-1; i++ {
		var curr int = prices[i]
		var next int = prices[i+1]

		// valid new profit of adjacent days which is better than current profit
		if next-curr > profit && next-curr > next-buy {
			profit = next - curr
			buy = curr
			sell = next
			buy_idx = i
			sell_idx = i + 1
		}

		// always find the lowest value as a potential buy value
		if curr < buy {
			low = curr
			low_idx = i
		}

		// check if the new sell value is a better sell option
		if next-buy > profit {
			sell = next
			sell_idx = i + 1
			profit = next - buy
		}

		// check if the new low option is valid and more profitable - sell value
		if sell-low > profit && low_idx < sell_idx {
			buy = low
			buy_idx = i
			profit = sell - buy
		}

		// check if the new low option is valid and more profitable - next value
		if next-low > profit {
			buy = low
			buy_idx = i
			sell = next
			sell_idx = i + 1
			profit = next - buy
		}
	}

	fmt.Println(buy_idx)
	if buy < sell {
		return sell - buy
	} else {
		return 0
	}
}
