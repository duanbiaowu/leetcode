package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

// 穿针引线: 头插法[画图清晰表明过程]
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}
