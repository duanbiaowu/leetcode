package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type Node = structures.RandomListNode

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 复制节点跟在原节点后面
	// 1->2->3  ==>  1->1'->2->2'->3->3'
	cur := head
	for cur != nil {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next.Next
	}

	// 处理random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 分离两个链表
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		tmp := cur.Next
		cur.Next = cur.Next.Next
		cur = tmp
	}

	return cloneHead
}
