package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	if head.Val == next.Val {
		for next != nil && head.Val == next.Val {
			next = next.Next
		}
		head = deleteDuplicates(next)
	} else {
		head.Next = deleteDuplicates(next)
	}

	return head
}

func deleteDuplicatesIteratively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
