package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
)

type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorIteratively(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	ancestor := root
	for {
		if p.Val < ancestor.Val && q.Val < ancestor.Val {
			ancestor = ancestor.Left
		} else if p.Val > ancestor.Val && q.Val > ancestor.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}
}

func lowestCommonAncestorUsingPath(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	pathP := getPath(root, p)
	pathQ := getPath(root, q)

	var ancestor *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return ancestor
}

func getPath(root, target *TreeNode) []*TreeNode {
	var path []*TreeNode
	node := root
	for node != nil && node != target {
		path = append(path, node)
		if node.Val > target.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// target 节点也加入path
	// 可能存在: p是q父节点 || q是p 父节点的情况
	path = append(path, node)
	return path
}
