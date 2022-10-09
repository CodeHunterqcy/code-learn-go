package mid

/*
https://leetcode.cn/problems/word-search/?favorite=2cktkvj

给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

*/

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var help func(int, int, int) bool
	help = func(i int, j int, k int) bool {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}

		temp := board[i][j]
		board[i][j] = '/'
		res := help(i+1, j, k+1) || help(i-1, j, k+1) || help(i, j+1, k+1) || help(i, j-1, k+1)
		board[i][j] = temp
		return res
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if help(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
