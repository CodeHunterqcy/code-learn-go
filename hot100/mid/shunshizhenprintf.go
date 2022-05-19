package mid

/*
顺时针打印数组
*/
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	var res []int
	// 定义上下左右
	top, button, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= button && left <= right {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top += 1
		for i := top; i <= button; i++ {
			res = append(res, matrix[i][right])
		}
		right -= 1
		if top > button || left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[button][i])
		}
		button -= 1
		for i := button; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left += 1
	}
	return res
}
