package easy

// 只出现一次的数字，别的数字出现两次
func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		//这个题目不要额外的空间，考异或，记住就行了
		res = res ^ num
	}
	return res
}
