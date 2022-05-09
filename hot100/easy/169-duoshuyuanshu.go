package easy

import "sort"

//给定一个大小为 n 的数组nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 n/2 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。

//链接：https://leetcode.cn/problems/majority-element
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
	// 要求时间的复杂度On
}

//摩尔投票法
func majorityElement11(nums []int) int {
	res, count := 0, 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if res == nums[i] {
			count++
		} else {
			count--
		}

	}
	return res
}
