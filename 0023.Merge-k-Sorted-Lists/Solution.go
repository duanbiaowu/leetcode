package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// simple solution
//func mergeKLists(lists []*ListNode) *ListNode {
//	dummy := &ListNode{0, nil}
//	cur := dummy
//	for _, node := range lists {
//		cur = mergeTwoLists(cur, node)
//	}
//	return dummy.Next
//}

// divide-and-merge solution
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := left + (right-left)>>1
	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
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
