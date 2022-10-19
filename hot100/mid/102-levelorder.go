package mid

/*
https://leetcode.cn/problems/binary-tree-level-order-traversal/?favorite=2cktkvj
*/

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		size := len(nodes)
		var values []int
		for i := 0; i < size; i++ {
			values = append(values, nodes[i].Val)
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
		}

		res = append(res, values)
		nodes = nodes[size:]
	}
	return res
}

// 深度优先遍历
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) <= depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 0)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
