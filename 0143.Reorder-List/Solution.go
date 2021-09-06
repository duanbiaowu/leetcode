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
