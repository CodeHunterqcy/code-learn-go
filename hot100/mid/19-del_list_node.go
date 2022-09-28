package mid

/*

https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?favorite=2cktkvj
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	firstNode := head
	for n-1 > 0 && firstNode.Next != nil {
		firstNode = firstNode.Next
		n--
	}
	cur := dummyHead
	for firstNode.Next != nil {
		cur = cur.Next
		firstNode = firstNode.Next
	}

	cur.Next = cur.Next.Next

	return dummyHead.Next
}

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	firstNode, cur := head, dummyHead
	for i := 0; i < n-1; i++ {
		firstNode = firstNode.Next
	}
	for firstNode.Next != nil {
		firstNode = firstNode.Next
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}
