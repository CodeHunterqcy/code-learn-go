package hard

/*
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。



示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：

输入：s = ""
输出：0

https://leetcode.cn/problems/longest-valid-parentheses/solution/zui-chang-you-xiao-gua-hao-by-leetcode-solution/
*/

// longestValidParenthesesNO 暴力递归，以此判断每一段字符串是否合法-此方法超时
func longestValidParenthesesNo(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if (j-i+1)%2 == 1 {
				continue
			}
			if isValid(s[i : j+1]) {
				if j-i+1 > res {
					res = j - i + 1
				}
			}
		}
	}
	return res
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	mapKey := map[byte]byte{
		')': '(',
	}

	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && mapKey[s[i]] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

// 方法一：辅助栈
func longestValidParentheses1(s string) int {
	maxRes := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxRes = getMax(maxRes, i-stack[len(stack)-1])
			}
		}
	}

	return maxRes
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
