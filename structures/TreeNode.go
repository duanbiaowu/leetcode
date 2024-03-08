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

type BST struct {
	root *TreeNode
}

// Insert 插入一个元素.
func (bst *BST) Insert(value int) {
	newNode := &TreeNode{value, nil, nil}
	// 如果二叉树为空，那么这个节点就当作跟节点
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

// 从根节点依次比较
func insertNode(root, newNode *TreeNode) {
	if newNode.Val < root.Val { // 应该放到根节点的左边
		if root.Left == nil {
			root.Left = newNode
		} else {
			insertNode(root.Left, newNode)
		}
	} else if newNode.Val > root.Val { // 应该放到根节点的右边
		if root.Right == nil {
			root.Right = newNode
		} else {
			insertNode(root.Right, newNode)
		}
	}
	// 否则等于根节点
}

// Remove 删除一个元素.
func (bst *BST) Remove(value int) bool {
	_, existed := remove(bst.root, value)
	return existed
}

// 用来递归移除节点的辅助方法.
// 返回替换root的新节点，以及元素是否存在
func remove(root *TreeNode, value int) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	var existed bool
	// 从左边找
	if value < root.Val {
		root.Left, existed = remove(root.Left, value)
		return root, existed
	}
	// 从右边找
	if value > root.Val {
		root.Right, existed = remove(root.Right, value)
		return root, existed
	}
	// 如果此节点正是要移除的节点,那么返回此节点，同时返回之前可能需要调整.
	existed = true
	// 如果此节点没有孩子，直接返回即可
	if root.Left == nil && root.Right == nil {
		root = nil
		return root, existed
	}
	// 如果左子节点为空, 提升右子节点
	if root.Left == nil {
		root = root.Right
		return root, existed
	}
	// 如果右子节点为空, 提升左子节点
	if root.Right == nil {
		root = root.Left
		return root, existed
	}
	// 如果左右节点都存在,那么从右边节点找到一个最小的节点提升，这个节点肯定比左子树所有节点都大.
	// 也可以从左子树节点中找一个最大的提升，道理一样.
	smallestInRight, _ := min(root.Right)
	// 提升
	root.Val = smallestInRight
	// 从右边子树中移除此节点
	root.Right, _ = remove(root.Right, smallestInRight)
	return root, existed
}

// Search 搜索元素(检查元素是否存在)
func (bst *BST) Search(value int) bool {
	return search(bst.root, value)
}
func search(n *TreeNode, value int) bool {
	if n == nil {
		return false
	}
	if value < n.Val {
		return search(n.Left, value)
	}
	if value > n.Val {
		return search(n.Right, value)
	}
	return true
}

// Min 二叉搜索树中的最小值
func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}
func min(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	// 从左边找
	for {
		if n.Left == nil {
			return n.Val, true
		}
		n = n.Left
	}
}

// Max 二叉搜索树中的最大值
func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}
func max(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	// 从右边找
	for {
		if n.Right == nil {
			return n.Val, true
		}
		n = n.Right
	}
}

// PreOrderTraverse 前序遍历
func (bst *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(bst.root, f)
}
func preOrderTraverse(n *TreeNode, f func(int)) {
	if n != nil {
		f(n.Val) // 前
		preOrderTraverse(n.Left, f)
		preOrderTraverse(n.Right, f)
	}
}

// PostOrderTraverse 后序遍历
func (bst *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(bst.root, f)
}
func postOrderTraverse(n *TreeNode, f func(int)) {
	if n != nil {
		postOrderTraverse(n.Left, f)
		postOrderTraverse(n.Right, f)
		f(n.Val) // 后
	}
}
