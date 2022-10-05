package mid

/*
https://leetcode.cn/problems/rotate-image/?favorite=2cktkvj
*/
// rotate 想得到顺时针旋转的数组，先将数组行与列交换后，再进行一次垂直对称旋转，就得到了。
func rotate(matrix [][]int) {
	// 先行变列
	len := len(matrix)
	for i := 0; i < len; i++ {
		for j := i; j < len; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 再对称反转
	for i := 0; i < len; i++ {
		for j := 0; j < len/2; j++ {
			matrix[i][j], matrix[i][len-j-1] = matrix[i][len-j-1], matrix[i][j]
		}
	}
}
