package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 快指针先走一步，避免链表只有一个节点时误判为环形链表
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return false
}
