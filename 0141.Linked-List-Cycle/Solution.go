package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func hasCycle(head *ListNode) bool {
	// 需要注意的是，即使只有一个节点，也可能是环形链表
	// 所以这段代码无法通过测试
	// if head == nil || head.Next == nil {
	// 	return false
	// }
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
