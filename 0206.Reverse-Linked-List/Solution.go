package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

func reverseListForGolang(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}

	return prev
}
