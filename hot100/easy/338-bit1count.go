package easy

/*
*给你一个整数 n ，对于 0 <= i <= n 中的每个 i ，计算其二进制表示中 1 的个数 ，返回一个长度为 n + 1 的数组 ans 作为答案。
 */
func countBits(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		res = append(res, getNum1Count(i))
	}
	return res
}
func getNum1Count(num int) int {
	res := 0

	// 判断二进制有几个1
	for num != 0 {
		num &= num - 1
		res++
	}

	return res
}
