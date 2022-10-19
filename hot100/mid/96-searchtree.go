package mid

/*

96. 不同的二叉搜索树

https://leetcode.cn/problems/unique-binary-search-trees/?favorite=2cktkvj

dp[i]代表i个节点构成BST共有几种
dp[1]*dp[i-1]+dp[2]*dp[i-2]+dp[3]*dp[i-3]+~~~~~+dp[i]dp[1]
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i-1; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}

	return dp[n]
}
