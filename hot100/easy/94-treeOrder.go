package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归版本
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var help func(root *TreeNode)
	help = func(root *TreeNode) {
		if root.Left != nil {
			help(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			help(root.Right)
		}
	}
	help(root)
	return res
}

// 迭代版中序遍历
func inorderTraversal1(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}

	return res
}

// 迭代版中序遍历2
func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

//  前序遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		root = node.Right
	}

	return res
}

// 后序遍历 左右根=根右左的反转
func postorderTraversal(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Left
	}
	left, right := 0, len(res)-1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}
