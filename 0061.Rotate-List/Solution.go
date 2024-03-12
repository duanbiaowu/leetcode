package leetcode

import "github.com/duanbiaowu/leetcode/structures"

// ListNode structures.ListNode
type ListNode = structures.ListNode

// 所谓链表的旋转，就是链表的拆分再合并问题。
// 计算出链表的长度 L, 向右移动 K 个位置，就是将第 K 到第 N 个节点作为链表的前半部分，后半部分就是第 0 到第 K - 1 个节点，问题梳理清楚之后，代理实现就可以一气呵成了。
// 1. 计算链表的长度
// 2. 计算旋转偏移量
// 3. 将链表首位相连，形成循环链表
// 4. 将链表以偏移量进行切割

// 这里有个优化点: 通过链表长度 L 取模移动位置 K 得到需要旋转的偏移量，当链表长度减去偏移量等于 0 时，说明链表不需要发生任何变化，直接返回即可。

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
