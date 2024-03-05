package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

// 穿针引线: 头插法[画图清晰表明过程]
// 起始状态: 1 -> 2 -> 3 -> 4 -> 5，反转中间 3 个元素 2 -> 3 -> 4
// 最终结果: 1 -> 4 -> 3 -> 2 -> 5

// 代码执行分解

// prev = 1
// cur  = 2

// 第一轮循环
// next = 3
// cur.next = 4		 2 -> 4
// next.next = 2     3 -> 2 -> 4
// prev.next = 3     1 -> 3 -> 2 -> 4 -> 5
// 第一轮循环结果: 	  1 -> 3 -> 2 -> 4 -> 5

// 第二轮循环
// next = 4
// cur.next = 5		 2 -> 5
// next.next = 2     4 -> 3 -> 2
// prev.next = 4     1 -> 4 -> 3 -> 2 -> 5
// 第二轮循环结果: 	  1 -> 4 -> 3 -> 2 -> 5

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
