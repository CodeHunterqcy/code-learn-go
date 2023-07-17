package mid

/*
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 原地修改
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{}
	dummy.Next = head
	firstPre, curPre := dummy, dummy
	cur := head
	for cur != nil {
		if cur.Val < x {
			if firstPre.Next == cur {
				curPre = cur
				firstPre = cur
				cur = cur.Next
			} else {
				curPre.Next = cur.Next

				cur.Next = firstPre.Next
				firstPre.Next = cur
				firstPre = cur
				cur = curPre.Next
			}
		} else {
			curPre = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}

// 开辟两条链拼接
func partition1(head *ListNode, x int) *ListNode {
	smallList := &ListNode{}
	smallHead := smallList
	bigList := &ListNode{}
	bigHead := bigList

	cur := head
	for cur != nil {
		if cur.Val < x {
			smallList.Next = cur
			smallList = smallList.Next
		} else {
			bigList.Next = cur
			bigList = bigList.Next
		}
		cur = cur.Next
	}

	bigList.Next = nil
	smallList.Next = bigHead
	return smallHead.Next
}
