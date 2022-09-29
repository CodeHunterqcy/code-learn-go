package mid

/*
数字 n代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。


示例 1：

输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：

输入：n = 1
输出：["()"]
https://leetcode.cn/problems/generate-parentheses/?favorite=2cktkvj
*/

/*
深度优先搜+剪枝
当前左右括号都有大于 00 个可以使用的时候，才产生分支；
产生左分支的时候，只看当前是否还有左括号可以使用；
产生右分支的时候，还受到左分支的限制，右边剩余可以使用的括号数量一定得在严格大于左边剩余的数量的时候，才可以产生分支；
在左边和右边剩余的括号数都等于 00 的时候结算。
*/
func generateParenthesis(n int) []string {
	var res []string
	var dfs func(int, int, string)
	dfs = func(left int, right int, path string) {
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			dfs(left-1, right, path+"(")
		}
		if right > 0 {
			dfs(left, right-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

func generateParenthesis1(n int) []string {
	var res []string
	var dfs func(int, int, string)
	dfs = func(left int, right int, path string) {
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}
		if left > right {
			return
		}
		if left > 0 {
			dfs(left-1, right, path+"(")
		}
		if right > 0 {
			dfs(left, right-1, path+")")
		}
	}
	dfs(n, n, "")

	return res
}
