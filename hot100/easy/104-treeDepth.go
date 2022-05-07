package easy

// 二叉树最大深度
func maxDepth(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	}

	return GetMax(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

//
func maxDepth1(root *TreeNode) int {

}
