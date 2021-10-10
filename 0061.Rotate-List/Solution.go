package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 计算链表长度
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}

	// 计算旋转偏移量
	offset := n - k%n
	if offset == 0 {
		return head
	}

	// 链表成环形
	cur.Next = head
	for offset > 0 {
		cur = cur.Next
		offset--
	}

	// 链表切断环形
	res := cur.Next
	cur.Next = nil

	return res
}
