package mid

/*
无重复字符的最长子串
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	res := 0
	if len(s) == 0 {
		return res
	}
	if len(s) == 1 {
		return 1
	}
	left, right := 0, 0
	set := make(map[byte]bool)
	for left <= right && right < len(s) {
		if _, ok := set[s[right]]; !ok {
			set[s[right]] = true
			right++
		} else {
			delete(set, s[left])
			left++
		}
		length := right - left
		if length > res {
			res = length
		}
	}
	return res
}
