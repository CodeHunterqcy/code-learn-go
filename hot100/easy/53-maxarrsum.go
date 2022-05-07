package easy

import "math"

// maxSubArray 最大子数组和贪心版本
func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}

	res, sum := math.MinInt64, 0
	for i := 0; i < length; i++ {
		sum += nums[i]
		if sum > res {
			res = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}

	return res
}

// maxSubArrayDP 动态规划版本
func maxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// dp【i】是截止到i的最大子序列和
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = GetMax(nums[i]+dp[i-1], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
