package leetocde

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head
	return next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	cur := dummy
	// cur 必须从 dummy 开始
	// 这样才可以更换 dummy.Next (head) 指针的指向
	// 因为最终要返回 dummy.Next

	// cur 初始化时指向 dummy
	// 但是在遍历过程中会不断发生变更
	// 而 dummy 只会在第一次指向 head.Next.Next 之后就不会发生改变
	// 正确的答案，最终的 dummy.Next 指向的是源链表的第二个节点

	// Example:
	// 1 -> 2 -> 3 -> 4

	// 初始化
	// dummy -> 1 -> 2 -> 3 -> 4

	// 开始执行循环
	// dummy -> 2 -> 1 -> 3 -> 4
	// dummy -> 2 -> 1 -> 4 -> 3
	// 结束循环
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next

		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		cur = node1
	}

	return dummy.Next
}
