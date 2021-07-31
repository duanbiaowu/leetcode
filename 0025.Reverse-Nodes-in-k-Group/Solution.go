package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 当链表只有一个节点的时候直接返回，否则会形成 "环"
	if k <= 0 || head == nil || head.Next == nil {
		return head
	}

	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	dummy := &ListNode{Next: head}
	dummy.Next = reverse(head, cur)
	head.Next = reverseKGroup(cur, k)
	return dummy.Next
}

func reverse(head, tail *ListNode) *ListNode {
	// 1->2->3 		head:1 tail:2
	// return 2->1 	head:2 tail: 1
	// 有可能翻转链表中间某一段，所以 prev := tail
	// 最终 prev 指向 tail 前面一个 node
	var prev *ListNode
	for head != tail {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}
