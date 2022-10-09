package mid

/*
https://leetcode.cn/problems/subsets/

https://leetcode.cn/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
*/
func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(i int, path []int) {

		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		for j := i; j < len(nums); j++ {
			path = append(path, nums[j])
			dfs(j+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return res
}

func subsets1(nums []int) [][]int {
	var res [][]int
	var dfs func(int, []int)
	dfs = func(i int, path []int) {
		if i == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		path = append(path, nums[i])
		dfs(i+1, path)
		path = path[:len(path)-1]
		dfs(i+1, path)
	}
	dfs(0, []int{})
	return res
}
