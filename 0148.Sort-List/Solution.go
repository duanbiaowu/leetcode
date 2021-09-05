package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur, p1, p2 := dummy, head1, head2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			cur.Next = p2
			p2 = p2.Next
		} else {
			cur.Next = p1
			p1 = p1.Next
		}
		cur = cur.Next
	}

	if p1 != nil {
		cur.Next = p1
	} else if p2 != nil {
		cur.Next = p2
	}

	return dummy.Next
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow
	return merge(sort(head, mid), sort(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}
