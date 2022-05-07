package easy

// 判断是不是对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var help func(left *TreeNode, right *TreeNode) bool
	help = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}

		return left.Val == right.Val && help(left.Right, right.Left) && help(left.Left, right.Right)
	}

	return help(root.Left, root.Right)
}
