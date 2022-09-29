package mid

import "sort"

/*
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

示例 1：

输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：

输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))
	var dfs func([]int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			dfs(path)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(make([]int, len(res)))
	return res
}
