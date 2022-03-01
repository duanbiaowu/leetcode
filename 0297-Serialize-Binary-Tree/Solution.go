package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"strconv"
	"strings"
)

type TreeNode = structures.TreeNode

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// 先序遍历
// Serializes a tree to a single string.
func (this *Codec) serialize2(root *TreeNode) string {
	var sb strings.Builder
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize2(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}

// 层序遍历 BFS
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	var sb strings.Builder
	var queue []*TreeNode

	queue = append(queue, root)
	sb.WriteByte('[')

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			sb.WriteString(strconv.Itoa(node.Val))
			sb.WriteByte(',')
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			sb.WriteString("null,")
		}
	}

	s := sb.String()
	// remove last , character
	s = s[:len(s)-1]
	return s + "]"
}

func (this *Codec) deserialize(data string) *TreeNode {
	n := len(data)
	if n < 2 || data == "[]" {
		return nil
	}

	values := strings.Split(data[1:n-1], ",")
	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: val}

	var queue []*TreeNode
	queue = append(queue, root)
	index := 1
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if values[index] != "null" {
			val, _ = strconv.Atoi(values[index])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		index++

		if values[index] != "null" {
			val, _ = strconv.Atoi(values[index])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		index++
	}

	return root
}
