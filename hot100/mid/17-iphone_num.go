package mid

/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	numMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string

	var dfs func(int, string)
	dfs = func(n int, path string) {
		if n == len(digits) {
			res = append(res, path)
			return
		}
		for _, val := range numMap[string(digits[n])] {
			dfs(n+1, path+string(val))
		}
	}

	dfs(0, "")
	return res
}

func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var mapping = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	var res []string
	var dfs func(int, []byte)
	dfs = func(n int, path []byte) {
		if n == len(digits) {
			res = append(res, string(path))
			return
		}
		for _, val := range mapping[digits[n]] {
			dfs(n+1, append(path, val))
		}
	}

	dfs(0, []byte{})
	return res
}
