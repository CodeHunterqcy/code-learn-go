package mid

/*
https://leetcode.cn/problems/sort-colors/?favorite=2cktkvj
*/
func sortColors(nums []int) {

}

// 双指针,用p0交换0，用p1交换1
func sortColors1(nums []int) {
	p0, p1 := 0, 0
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if num == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}
