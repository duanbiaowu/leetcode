package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for level := 0; len(queue) > 0; level++ {
		n := len(queue)
		res = append(res, make([]int, n))
		for i := 0; i < n; i++ {
			res[level][i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]

		// 偶数层级反转数组
		if level&1 == 1 {
			for left, right := 0, n-1; left < right; {
				res[level][left], res[level][right] = res[level][right], res[level][left]
				left++
				right--
			}
		}
	}

	return res
}

func levelOrder(root *TreeNode) [][]int {
	var dfs func(root *TreeNode, depth int, res *[][]int)
	dfs = func(root *TreeNode, depth int, res *[][]int) {
		if root == nil {
			return
		}
		if depth == len(*res) {
			*res = append(*res, []int{})
		}
		if depth&1 == 1 {
			(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
		} else {
			(*res)[depth] = append((*res)[depth], root.Val)
		}
		dfs(root.Left, depth+1, res)
		dfs(root.Right, depth+1, res)
	}

	var res [][]int
	dfs(root, 0, &res)
	return res
}
