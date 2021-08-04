package structures

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = math.MinInt64

func GenerateTreeNodesBySlice(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}

	queue := make([]*TreeNode, 1, n/2+1)
	queue[0] = root

	for top, index := 0, 1; index < n; index++ {
		node := queue[top]
		top++

		if nums[index] != NULL {
			node.Left = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Left)
		}
		index++

		if index < n && nums[index] != NULL {
			node.Right = &TreeNode{Val: nums[index]}
			queue = append(queue, node.Right)
		}
	}

	return root
}
