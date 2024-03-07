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
	// 这里哨兵节点极大提高了代码可读性
	// 直接从 cur.Next (也就是头节点) 开始
	// 这样即时头节点就是重复元素，也可以直接将其删除，然后将哨兵节点的 Next 指针指向到第一个不重复的节点
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			// 删除掉所有和当前节点值相等的后续节点
			// 注意这里不改变 cur 指针的指向，而是直接改变 cur.Next 指针的指向
			// 也就是说，每次判断的都是以 [下个节点] 和 [下下个节点]
			// 当前节点从哨兵节点开始，本身就是不重复的值，因为重复的值在节点连接过程中都被删除掉了
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
