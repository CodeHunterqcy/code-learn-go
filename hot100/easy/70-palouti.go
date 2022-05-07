package easy

// 爬楼梯 n台阶楼梯，每次能爬1或者2，多少种方法
func climbStairs(n int) int {
	// 动态规划，分为两个子问题，爬n台阶等于爬n-1和爬n-2的和
	// dp[n] = dp[n-1]+dp[n-2]
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
