package mid

/*
最长回文子串
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}

	res := s[0:1]
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for l := 2; l <= length; l++ {
		for start := 0; start < length; start++ {
			end := start + l - 1
			if end >= length {
				break
			}
			if s[end] != s[start] {
				dp[start][end] = false
			} else {
				if end-start < 3 {
					dp[start][end] = true
				} else {
					dp[start][end] = dp[start+1][end-1]
				}
			}

			if dp[start][end] == true && end-start+1 > len(res) {
				res = s[start : end+1]
			}
		}
	}

	return res
}

// longestPalindrome1 动态规划版
func longestPalindrome1(s string) string {

	length := len(s)
	if length < 2 {
		return s
	}

	dp := make([][]bool, length)
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
		dp[i][i] = true
	}

	for l := 2; l <= length; l++ {
		for start := 0; start < length; start++ {
			end := start + l - 1
			if end >= length {
				break
			}

			if s[start] != s[end] { //首尾不同则不可能为回文
				dp[start][end] = false
			} else {
				if end-start < 3 {
					dp[start][end] = true //即case 2的判断
				} else {
					dp[start][end] = dp[start+1][end-1] //状态转移
				}
			}

			if dp[start][end] && end-start+1 > len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}

	return result
}
