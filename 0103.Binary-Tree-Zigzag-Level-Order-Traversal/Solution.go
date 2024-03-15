package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for level := 1; len(queue) > 0; level++ {
		length := len(queue)
		row := make([]int, length)

		for i := 0; i < length; i++ {
			row[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		// 偶数层反转数组
		if level&1 == 0 {
			reverse(row)
		}

		queue = queue[length:]

		res = append(res, row)
	}

	return res
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func zigzagLevelOrderDFS(root *TreeNode) [][]int {
	var res [][]int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}

	if depth == len(*res) {
		*res = append(*res, []int{})
	}

	// 因为 depth 变量除了表示层级外，还充当了返回值数组的索引
	// 所以起始是从 0 开始的
	// 当 depth % 2 == 1 时，表示当前层为偶数层
	if depth&1 == 1 {
		// 偶数层从末尾插入元素
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	} else {
		// 奇数层顺序插入元素
		(*res)[depth] = append((*res)[depth], root.Val)
	}

	dfs(root.Left, depth+1, res)
	dfs(root.Right, depth+1, res)
}
