package hard

/*
23. 合并K个升序链表
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

https://leetcode.cn/problems/merge-k-sorted-lists/?favorite=2cktkvj
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeKLists 遍历，分别俩俩合并
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	preList, cur := &ListNode{}, &ListNode{}
	for i, list := range lists {
		if i == 0 {
			preList = list
			continue
		}
		cur = list
		preList = mergeTwoList(preList, cur)
	}

	return preList
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummyHead := &ListNode{}
	cur := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return dummyHead.Next
}
