package mid

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

https://leetcode.cn/problems/container-with-most-water/?favorite=2cktkvj
*/

// maxArea 暴力递归-超时
func maxArea(height []int) int {
	// maxRes =Max(maxRes,Min(s[left],s[right])*(right-left))

	return 0
}

// maxArea1 双指针
func maxArea1(height []int) int {
	left, right := 0, len(height)-1
	maxRes := 0
	for left < right {
		maxRes = getMax(maxRes, getMin(height[left], height[right]*(right-left)))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxRes
}
