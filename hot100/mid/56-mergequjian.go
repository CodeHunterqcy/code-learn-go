package mid

import "sort"

/*
56. 合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

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

func merge1(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	pre := intervals[0]
	for _, interval := range intervals {
		if interval[0] <= pre[1] {
			pre = []int{pre[0], getMax(pre[1], interval[1])}
		} else {
			res = append(res, pre)
			pre = interval
		}
	}

	res = append(res, pre)
	return res
}
