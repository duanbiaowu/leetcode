package leetcode

import "math"

func verifyPostorder(postorder []int) bool {
	return dfs(postorder, 0, len(postorder)-1)
}

// 根据后序遍历规则，postorder[j] 即根节点
// 左子树区间 [i, m-1] 所有元素都应 < postorder[j]
// 右子树区间 [m, j-1] 所有元素都应 > postorder[j]
func dfs(postorder []int, i, j int) bool {
	if i >= j {
		return true
	}
	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	// 第一个大于根节点的节点
	m := p
	// 逐一判断右区间 所有元素都应 > postorder[j]
	for postorder[p] > postorder[j] {
		p++
	}
	return p == j && dfs(postorder, i, m-1) && dfs(postorder, m, j-1)
}

func verifyPostorder2(postorder []int) bool {
	var stack []int
	root := math.MaxInt
	for i := len(postorder) - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) > 0 && stack[len(stack)-1] > postorder[i] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
