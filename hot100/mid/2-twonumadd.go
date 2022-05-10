package mid

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表的两数相加
/*
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 注意进位
	plus := 0
	var head, cur *ListNode
	for l1 != nil || l2 != nil {
		l1Val, l2Val := 0, 0
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}
		sum := l1Val + l2Val + plus
		sum, plus = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			cur = head
		} else {
			cur.Next = &ListNode{Val: sum}
			cur = cur.Next
		}
	}
	if plus > 0 {
		cur.Next = &ListNode{Val: plus}
	}
	return head
}
