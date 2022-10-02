package mid

import "sort"

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：

输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：

输入：nums = [], target = 0
输出：[-1,-1]

要求时间复杂度 O(log n)
*/
// 时间复杂度不满足，如果数组全是target，时间复杂度退化为O（n）
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			start := mid - 1
			for start >= 0 && nums[start] == target {
				start--
			}
			end := mid + 1
			for end <= len(nums)-1 && nums[end] == target {
				end++
			}

			return []int{start + 1, end - 1}
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return []int{-1, -1}
}

// 答案版,用了sort里面的api
func searchRange1(nums []int, target int) []int {
	leftmost := sort.SearchInts(nums, target)
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}
	rightmost := sort.SearchInts(nums, target+1) - 1
	return []int{leftmost, rightmost}
}

// 答案版，手写
func searchRange2(nums []int, target int) []int {
	var searchIndex func(bool) int
	searchIndex = func(flag bool) int {
		left, right, res := 0, len(nums)-1, len(nums)
		for left <= right {
			mid := (right-left)/2 + left
			if nums[mid] > target || flag && nums[mid] >= target {
				right = mid - 1
				res = mid
			} else {
				left = mid + 1
			}
		}
		return res
	}

	leftIndex := searchIndex(true)
	rightIndex := searchIndex(false) - 1
	if leftIndex <= rightIndex && rightIndex < len(nums) && nums[leftIndex] == target && nums[rightIndex] == target {
		return []int{leftIndex, rightIndex}
	}

	return []int{-1, -1}
}
