package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	var queue = []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, queue[n-1].Val)
		queue = queue[n:]
	}

	return res
}

// 前序遍历反向操作: 根结点 -> 右子树 -> 左子树
// 保证每层都是最先访问最右边的节点
func rightSideViewDFS(root *TreeNode) []int {
	var res []int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *[]int) {
	if root == nil {
		return
	}
	if depth == len(*res) {
		*res = append(*res, root.Val)
	}
	dfs(root.Right, depth+1, res)
	dfs(root.Left, depth+1, res)
}
