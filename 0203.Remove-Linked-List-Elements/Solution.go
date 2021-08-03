package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

func removeElementsRecursively(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	head.Next = removeElementsRecursively(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}
