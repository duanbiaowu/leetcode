package leetcode

import "leetcode/structures"

type ListNode = structures.ListNode

// simple solution
func mergeKListsSimple(lists []*ListNode) *ListNode {
	var head *ListNode

	for i := range lists {
		head = mergeTwoLists(head, lists[i])
	}

	return head
}

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

// priority-queue solution
//func mergeKLists(lists []*ListNode) *ListNode {
//	if len(lists) == 0 {
//		return nil
//	}
//
//	pq := make(ListNodeQueue, len(lists))
//	for i, node := range lists {
//      // 可以入队时判断是否为nil
//		pq[i] = node
//	}
//
//	heap.Init(&pq)
//	dummy := &ListNode{0, nil}
//	cur := dummy
//
//	for pq.Len() > 0 {
//		next := heap.Pop(&pq).(*ListNode)
// 		// 因为入队的时候没有判断是否为nil，所以出队后需要判断
//		if next != nil {
//			cur.Next = next
//			cur = cur.Next
//			if next.Next != nil {
//				heap.Push(&pq, next.Next)
//			}
//		}
//	}
//
//	return dummy.Next
//}

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
