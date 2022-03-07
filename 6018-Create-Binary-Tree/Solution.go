package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode)
	// Set 里面存放所有非 root 节点 key
	set := make(map[int]struct{})

	for i := range descriptions {
		parent := descriptions[i][0]
		child := descriptions[i][1]
		left := descriptions[i][2]

		if _, ok := m[parent]; !ok {
			m[parent] = &TreeNode{Val: parent}
		}
		if _, ok := m[child]; !ok {
			m[child] = &TreeNode{Val: child}
		}
		if left == 1 {
			m[parent].Left = m[child]
		} else {
			m[parent].Right = m[child]
		}

		set[child] = struct{}{}
	}

	for key := range m {
		if _, ok := set[key]; !ok {
			return m[key]
		}
	}
	return nil
}
