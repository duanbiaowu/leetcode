package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for _, node := range lists {
		cur = mergeTwoLists(cur, node)
	}
	return dummy.Next
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
