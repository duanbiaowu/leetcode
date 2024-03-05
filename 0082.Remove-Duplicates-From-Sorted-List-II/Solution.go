package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	if head.Val == next.Val {
		for next != nil && head.Val == next.Val {
			next = next.Next
		}
		head = deleteDuplicates(next)
	} else {
		head.Next = deleteDuplicates(next)
	}

	return head
}

// 注释版本 (提高可读性)
func deleteDuplicatesWithComment(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		// 当前节点重复
		next := head.Next
		// 先将当前节点和后面的重复节点全部删除
		for next != nil && head.Val == next.Val {
			next = next.Next
		}
		// 从下一个节点开始删除重复节点
		head = deleteDuplicates(next)
	} else {
		// 当前节点不重复
		// 跳过
		// 从下一个节点开始删除重复节点
		head.Next = deleteDuplicates(head.Next)
	}

	return head
}

func deleteDuplicatesIteratively(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
