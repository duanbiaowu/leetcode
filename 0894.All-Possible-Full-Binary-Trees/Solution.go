package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

var (
	memo = map[int][]*TreeNode{} // 避免重复计算
)

func allPossibleFBT(n int) []*TreeNode {
	var res []*TreeNode

	if n <= 0 || n%2 == 0 {
		return res
	}
	if val, ok := memo[n]; ok {
		return val
	}
	if n == 1 {
		res = append(res, &TreeNode{Val: 0})
		return res
	}

	for i := 1; i < n-1; i += 2 {
		leftTrees := allPossibleFBT(i)
		rightTress := allPossibleFBT(n - i - 1)

		memo[i] = leftTrees
		memo[n-i-1] = rightTress

		for left := range leftTrees {
			for right := range rightTress {
				root := &TreeNode{Val: 0}
				root.Left = leftTrees[left]
				root.Right = rightTress[right]
				res = append(res, root)
			}
		}
	}

	return res
}
