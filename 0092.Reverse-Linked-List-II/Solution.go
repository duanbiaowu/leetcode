package leetcode

import (
	"fmt"

	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

// 穿针引线: 头插法[画图清晰表明过程]
// 起始状态: 1 -> 2 -> 3 -> 4 -> 5，反转中间 3 个元素 2 -> 3 -> 4
// 最终结果: 1 -> 4 -> 3 -> 2 -> 5

// 整体思路:
// 在需要被反转的区间链表里，每遍历到一个节点时，让当前节点后面的节点插入到当前节点前面

// 代码执行分解

// prev = 1
// cur  = 2

// 第一轮循环，当前节点为 2
// cur 永远指向的是第 left 的位置，不论这个位置的节点值发生什么变化
// next 永远指向 cur 的下一个节点
// prev 永远指向 cur 的前一个节点

// next = 3
// cur.next = 4		 2 -> 4
// next.next = 2     3 -> 2 -> 4
// prev.next = 3     1 -> 3 -> 2 -> 4 -> 5
// 第一轮循环结果: 	  1 -> 3 -> 2 -> 4 -> 5

// 第二轮循环，当前节点为 3
// cur 永远指向的是第 left 的位置，不论这个位置的节点值发生什么变化
// next 永远指向 cur 的下一个节点
// prev 永远指向 cur 的前一个节点

// next = 4
// cur.next = 5		 2 -> 5
// next.next = 2     4 -> 3 -> 2
// prev.next = 4     1 -> 4 -> 3 -> 2 -> 5
// 第二轮循环结果: 	  1 -> 4 -> 3 -> 2 -> 5

// reference: https://leetcode.cn/problems/reverse-linked-list-ii/solutions/634701/fan-zhuan-lian-biao-ii-by-leetcode-solut-teyq/
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

func reverseBetweenSimple(head *ListNode, left int, right int) *ListNode {
	if head == nil || right < left {
		return head
	}

	dummy := &ListNode{Next: head}

	// 获取左边部分链表的最后一个节点
	leftLast := dummy
	for i := 0; i < left-1; i++ {
		leftLast = leftLast.Next
	}

	// 获取右边部分的链表的头节点的上一个节点
	// 也就是中间部分的链表的最后一个节点
	midLast := leftLast
	for i := 0; i < right-left+1; i++ {
		midLast = midLast.Next
	}

	// 获取中间部分的链表的第一个节点
	midHead := leftLast.Next
	rightHead := midLast.Next

	// 切换链表左右部分的连接，形成一个子链表
	leftLast.Next = nil
	midLast.Next = nil

	reverseList(midHead)

	// 反转之后
	// 左边部分链表应该连接中间部分链表的最后一个节点
	leftLast.Next = midLast
	// 中间部分链表应该连接右边部分链表的第一个节点
	midHead.Next = rightHead

	return dummy.Next
}

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
