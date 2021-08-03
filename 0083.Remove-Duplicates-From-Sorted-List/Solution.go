package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	// 问题转化为: 删除某个值 if head.Val == someVal
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}
