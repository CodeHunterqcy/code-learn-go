package easy

// 反转链表
func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
