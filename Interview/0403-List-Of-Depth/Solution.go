package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

type ListNode = structures.ListNode

func listOfDepth(tree *TreeNode) []*ListNode {
	if tree == nil {
		return nil
	}

	var res []*ListNode
	var queue []*TreeNode
	queue = append(queue, tree)
	for n := len(queue); n > 0; n = len(queue) {
		dummy := &ListNode{}
		cur := dummy
		for i := 0; i < n; i++ {
			cur.Next = &ListNode{Val: queue[i].Val}
			cur = cur.Next
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, dummy.Next)
		queue = queue[n:]
	}

	return res
}
