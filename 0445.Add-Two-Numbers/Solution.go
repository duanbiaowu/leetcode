package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverse(l1)
	l2 = reverse(l2)

	dummy := &ListNode{}

	n1, n2, carry := 0, 0, 0

	cur := dummy
	for l1 != nil || l2 != nil || carry > 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		cur = cur.Next
		carry = (n1 + n2 + carry) / 10
	}

	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
