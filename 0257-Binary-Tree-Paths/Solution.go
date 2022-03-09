package leeetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"strconv"
	"strings"
)

type TreeNode = structures.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var path []string
	dfs(root, path, &res)
	return res
}

func dfs(root *TreeNode, path []string, res *[]string) {
	if root == nil {
		return
	}
	path = append(path, strconv.Itoa(root.Val))
	if root.Left == nil && root.Right == nil {
		*res = append(*res, strings.Join(path, "->"))
	}
	dfs(root.Left, path, res)
	dfs(root.Right, path, res)
}
