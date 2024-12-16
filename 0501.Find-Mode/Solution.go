package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func findMode(root *TreeNode) []int {
	var res []int

	callback := update(&res)
	dfs(root, callback)

	return res
}

// 返回闭包
// 便于随时保存/修改 内部变量的变化状态
func update(res *[]int) func(x int) {
	var base, cnt, maxCnt int

	return func(x int) {
		if x == base {
			cnt++
		} else {
			base = x
			cnt = 1
		}

		// 众数可能存在多个
		if cnt == maxCnt {
			*res = append(*res, base)
		} else if cnt > maxCnt {
			maxCnt = cnt
			*res = []int{base}
		}
	}
}

func dfs(root *TreeNode, update func(x int)) {
	if root == nil {
		return
	}

	dfs(root.Left, update)
	update(root.Val)
	dfs(root.Right, update)
}
