package leetcode

import "github.com/duanbiaowu/leetcode/structures"

type TreeNode = structures.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)>>1
	root := &TreeNode{Val: nums[mid]}
	root.Left = build(nums, left, mid-1)
	root.Right = build(nums, mid+1, right)

	return root
}
