package mid

// 给你一个字符串 s ，请你统计并返回这个字符串中 回文子串 的数目。
//回文字符串 是正着读和倒过来读一样的字符串。
//子字符串 是字符串中的由连续字符组成的一个序列。
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/palindromic-substrings
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// countSubstrings 回文子串
/*
输入：s = "abc"
输出：3
解释：三个回文子串: "a", "b", "c"
*/
func countSubstrings(s string) int {

	return 0
}

func isHuiWen(s string) bool {
	start, end := 0, len(s)
	if start+1 == end {
		return true
	}
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}
