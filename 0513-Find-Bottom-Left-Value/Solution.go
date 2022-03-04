package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func findBottomLeftValue(root *TreeNode) int {
	return bfs(root)
	//max, res := -1, -1
	//dfs(root, 0, &max, &res)
	//return res
}

func dfs(root *TreeNode, depth int, max, res *int) {
	if root == nil {
		return
	}
	if depth > *max {
		*max = depth
		*res = root.Val
	}
	dfs(root.Left, depth+1, max, res)
	dfs(root.Right, depth+1, max, res)
}

// 最底层 最左边 节点的值（也可能在右边）
// 从右到左，非常巧妙
func bfs(root *TreeNode) int {
	if root == nil {
		return -1
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		root = queue[0]
		queue = queue[1:]
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	}
	return root.Val
}
