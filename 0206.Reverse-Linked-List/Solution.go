package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type ListNode = structures.ListNode

// 1 -> 2 -> 3 -> 4 -> 5
// 5 -> 4 -> 3 -> 2 -> 1
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	// 从前往后扫描，逐个改变节点指向
	// 四步顺序操作 (不可改变)
	for head != nil {
		next := head.Next // 更新下一个节点为当前节点的下一个节点
		head.Next = prev  // 开始反转，将当前节点的下个节点指向前一个节点
		prev = head       // 更新前一个节点为当前节点
		head = next       // 更新当前节点位下一个节点
	}

	return prev
}

func reverseListForGolang(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		prev, head, head.Next = head, head.Next, prev
	}

	return prev
}

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
