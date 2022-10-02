package hard

/*
https://leetcode.cn/problems/trapping-rain-water/?favorite=2cktkvj
*/

//
func trap(height []int) int {
	res := 0
	for i := 1; i < len(height)-1; i++ {
		maxLeft := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}

		maxRight := 0
		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}

		min := getMin(maxLeft, maxRight)
		if min > height[i] {
			res += min - height[i]
		}
	}

	return res
}
