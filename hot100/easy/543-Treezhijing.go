package easy

// 二叉树直径=左树最大深度+右树最大深度
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}
	var help func(root *TreeNode) int
	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := help(root.Left)
		rightDepth := help(root.Right)
		res = GetMax(leftDepth+rightDepth, res)
		return GetMax(leftDepth, rightDepth) + 1
	}
	help(root)
	return res
}
