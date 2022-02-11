package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func BSTSequences(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	var dq []*TreeNode
	var path []int
	dq = append(dq, root)

	backtrack(dq, path, &res)
	return res
}

func backtrack(dq []*TreeNode, path []int, res *[][]int) {
	size := len(dq)
	if size == 0 {
		*res = append(*res, append([]int(nil), path...))
		return
	}

	for i := 0; i < size; i++ {
		cur := dq[0]
		dq = dq[1:]
		path = append(path, cur.Val)

		if cur.Left != nil {
			dq = append(dq, cur.Left)
		}
		if cur.Right != nil {
			dq = append(dq, cur.Right)
		}

		backtrack(dq, path, res)

		dq = dq[:size-1]
		dq = append(dq, cur)
		path = path[:len(path)-1]
	}
}
