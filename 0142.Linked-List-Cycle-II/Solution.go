package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// ListNode Definition for singly-linked list.
type ListNode = structures.ListNode

// a: head 到环入口长度
// b: 环长度

// 1. f = 2s （快指针每次 2 步）
// 2. f = s + nb (相遇时，刚好多走了 n 圈）
// 推导出 s = nb

// head 到环入口长度: a + nb
//		先走 a 步到入口节点，之后每绕 1 圈环（ b 步）都会再次到入口节点
// slow 已经走了 nb，再走 a 步就是入环点

// 如何求出 a 的值？
// 从 head 开始，和 slow 指针一起走，相遇时刚好就是 a 步
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}

	return nil
}
