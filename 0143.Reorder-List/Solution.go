package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	left, right := 0, len(nodes)-1
	for left < right {
		nodes[left].Next = nodes[right]
		left++
		nodes[right].Next = nodes[left]
		right--
	}
	nodes[left].Next = nil
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}

func mergeList(head1, head2 *ListNode) {
	var tmp1, tmp2 *ListNode
	for head1 != nil && head2 != nil {
		tmp1 = head1.Next
		tmp2 = head2.Next

		head1.Next = head2
		head1 = tmp1

		head2.Next = tmp1
		head2 = tmp2
	}
}

func reorderList2(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleNode(head)
	head1 := head
	head2 := mid.Next
	mid.Next = nil

	head2 = reverseList(head2)
	mergeList(head1, head2)
}
