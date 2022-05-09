package easy

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	res := root
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root == nil {
			return
		}

		root.Left, root.Right = root.Right, root.Left

		help(root.Left)
		help(root.Right)

	}
	help(res)
	return res

}
