package mid

import "sort"

/*
输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

*/
func merge(intervals [][]int) [][]int {
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		// 当前的和前一个比有重合
		if cur[0] <= pre[1] {
			pre = []int{pre[0], getMax(pre[1], cur[1])}
		} else { // 无重合
			res = append(res, pre)
			pre = cur
		}
	}
	res = append(res, pre)
	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
