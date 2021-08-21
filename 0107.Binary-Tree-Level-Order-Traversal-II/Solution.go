package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	for n := len(queue); n > 0; n = len(queue) {
		cur := len(res)
		res = append(res, make([]int, n))
		for i := 0; i < n; i++ {
			res[cur][i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}

	for left, right := 0, len(res)-1; left < right; {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}

	return res
}

func levelOrderBottomDFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	p := 0
	dfs(root, 0, &res, &p)

	for left, right := 0, len(res)-1; left < right; {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

func dfs(root *TreeNode, depth int, res *[][]int, p *int) {
	if root == nil {
		return
	}
	if depth >= len(*res) {
		*res = append(*res, []int{})
	}
	(*res)[depth] = append((*res)[depth], root.Val)
	dfs(root.Left, depth+1, res, p)
	dfs(root.Right, depth+1, res, p)
}
