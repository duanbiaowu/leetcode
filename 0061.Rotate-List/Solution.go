package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 1. 计算链表长度
	n := 1
	cur := head
	for cur.Next != nil {
		n++
		cur = cur.Next
	}

	// 2. 计算旋转偏移量
	// 偏移量 = 链表长度 - 移动位置%链表长度
	offset := n - k%n
	if offset == 0 {
		return head
	}

	// 3. 将链表首位相连，形成循环链表
	cur.Next = head
	for offset > 0 {
		offset--
		cur = cur.Next
	}

	// 4. 将链表以偏移量进行切割
	next := cur.Next
	cur.Next = nil

	return next
}
