package easy

// 买卖股票的最佳时机-暴力超时版
func maxProfit(prices []int) int {
	maxRes := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] {
				res := prices[j] - prices[i]
				if res > maxRes {
					maxRes = res
				}
			}
		}
	}

	return maxRes
}

// 想了想后的版本
func maxProfitDP(prices []int) int {
	minDay := prices[0]
	maxRes := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > minDay {
			res := prices[i] - minDay
			if res > maxRes {
				maxRes = res
			}
		} else {
			minDay = prices[i]
		}
	}

	return maxRes
}
