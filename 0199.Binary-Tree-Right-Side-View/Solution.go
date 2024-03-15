package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		res = append(res, queue[length-1].Val)
		queue = queue[length:]
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

	// 先递归右子树
	dfs(root.Right, depth+1, res)
	// 再递归左子树
	dfs(root.Left, depth+1, res)
}
