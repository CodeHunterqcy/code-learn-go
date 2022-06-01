package mid

/*
输入：strs = ["flower","flow","flight"]
输出："fl"
最长公共前缀
*/
func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		length := getMin(len(res), len(strs[i]))
		index := 0
		for index < length && res[index] == strs[i][index] {
			index++
		}
		res = strs[i][:index]
		if len(res) == 0 {
			break
		}
	}
	return res
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
