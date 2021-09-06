package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type ListNode = structures.ListNode

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	lastSorted := head
	cur := head.Next

	for cur != nil {
		// 当前节点大于已排序最后一个节点，直接连接起来，进入下一个节点
		if lastSorted.Val <= cur.Val {
			lastSorted = lastSorted.Next
		} else {
			// 找到节点对应的位置，连接起来
			prev := dummy
			for prev.Next.Val <= cur.Val {
				prev = prev.Next
			}

			lastSorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = lastSorted.Next
	}

	return dummy.Next
}
