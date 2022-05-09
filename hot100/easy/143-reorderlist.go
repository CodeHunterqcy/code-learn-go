package easy

/*
给定一个单链表 L 的头节点 head ，单链表 L 表示为：

L0 → L1 → … → Ln - 1 → Ln
请将其重新排列后变为：

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

链接：https://leetcode.cn/problems/reorder-list

*/
func reorderList(head *ListNode) {
	// 找到中点，反转后再拼接
	if head == nil || head.Next == nil {
		return
	}
	mid := findMidNode(head)
	rightHead := reverseList(mid)
	mid.Next = nil
	leftHead := head
	mergeList(leftHead, rightHead)
}
func mergeList(left, right *ListNode) {
	curL, curR := left, right
	for curL != nil && curR != nil {
		curL.Next, curR.Next, curR, curL = curR, curL.Next, curR.Next, curL.Next
	}
}
func findMidNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode {
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
