package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head
	for pre := dummy; pre.Next != nil; {
		cur := pre.Next
		for cur.Next != nil {
			if pre.Next.Val == cur.Next.Val {
				cur.Next = cur.Next.Next
			} else {
				cur = cur.Next
			}
		}
		pre = pre.Next
	}

	return dummy.Next
}

func removeDuplicateNodes2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	set := make(map[int]struct{})
	dummy := &ListNode{}
	dummy.Next = head
	for pre := dummy; pre.Next != nil; {
		if _, ok := set[pre.Next.Val]; ok {
			pre.Next = pre.Next.Next
		} else {
			set[pre.Next.Val] = struct{}{}
			pre = pre.Next
		}
	}

	return dummy.Next
}
