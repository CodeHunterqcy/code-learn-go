package mid

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。


示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/group-anagrams
*/

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	helpMap := make(map[[26]int][]string)
	for _, str := range strs {
		conStr := convert(str)
		helpMap[conStr] = append(helpMap[conStr], str)
	}

	for _, val := range helpMap {
		res = append(res, val)
	}

	return res
}

func convert(s string) [26]int {
	var ret [26]int
	for _, char := range s {
		ret[char-'a']++
	}
	return ret
}
