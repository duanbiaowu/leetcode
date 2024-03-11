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
		cur.Next = &Node{
			Val:  cur.Val,
			Next: cur.Next,
		}
		cur = cur.Next.Next
	}

	// 处理 Random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 分离源节点和影子节点
	// 将影子节点作为最终结果返回
	cur = head
	clonedHead := cur.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		cur = next
	}

	return clonedHead
}
