package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists 迭代版本
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	node := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			node.Next = list2
			list2 = list2.Next
		} else {
			node.Next = list1
			list1 = list1.Next
		}
		node = node.Next
	}

	if list2 != nil {
		node.Next = list2
	}
	if list1 != nil {
		node.Next = list1
	}

	return head.Next
}

// mergeListDiGui 递归版本
func mergeListDiGui(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeListDiGui(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeListDiGui(list1, list2.Next)
		return list2
	}
}
