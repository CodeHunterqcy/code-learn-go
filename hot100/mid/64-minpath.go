package mid

/*
https://leetcode.cn/problems/minimum-path-sum/?favorite=2cktkvj
*/
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += getMin(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[m-1][n-1]
}

func minPathSum1(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 0; j < n; j++ {
			dp[j] = getMin(dp[j], dp[j-1]) + grid[i][j]
		}
	}

	return dp[n-1]
}
