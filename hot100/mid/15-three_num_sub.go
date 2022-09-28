package mid

import "sort"

/*
15. 三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func threeSum(nums []int) [][]int {
	length := len(nums)
	var res [][]int
	sort.Ints(nums)

	for i := 0; i < length-2; i++ {
		n := nums[i]
		if n > 0 {
			break
		}
		if i > 0 && n == nums[i-1] {
			continue
		}
		left := i + 1
		right := length - 1
		for left < right {
			n1, n2 := nums[left], nums[right]
			sum := n1 + n + n2
			if sum == 0 {
				res = append(res, []int{n, n1, n2})
				for left < right && nums[left] == n1 {
					left++
				}
				for left < right && nums[right] == n2 {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
