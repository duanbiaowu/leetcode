package leetcode

// QuadTreeNode Definition for a QuadTree node.
type QuadTreeNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadTreeNode
	TopRight    *QuadTreeNode
	BottomLeft  *QuadTreeNode
	BottomRight *QuadTreeNode
}

func construct(grid [][]int) *QuadTreeNode {
	return dfs(grid, 0, len(grid))
}

// 题目翻译:
// 将一个大区域划分为若干个小区域（要求每个小区域中的 val 全部一致）
// 每次都是将一个大区域划分为 4 份（直到小区域val全部一致，终止）

// 解题步骤:
// 按照声明中的 topLeft、topRight、bottomLeft、bottomRight 将一个大区域划分为 4 个小区域
// 使用 for 循环校验每个小区域中 val 是否全部一致
// 如果小区域不一致，则为非叶子结点，递归继续划分（此时 val 为任意 bool 值）
// 如果小区域全部一致，则为叶子节点, isLeaf = true, val = bool(左上角基准值), 同时四个指针全部置为 nil

// 参数说明:
// 单次递归传入的 rows 都是上层递归的 1/4
func dfs(rows [][]int, left, right int) *QuadTreeNode {
	for _, row := range rows {
		for _, v := range row[left:right] {
			// 以当前检测区域的左上角小块为基准
			// 来判断当前区域是否所有值都一样
			if v != rows[0][left] {
				// 不是叶子节点
				// 将当前区域分割为 4 个子区域，递归进行切割
				rowMid, colMid := len(rows)>>1, (left+right)>>1
				return &QuadTreeNode{
					Val:         true,
					IsLeaf:      false,
					TopLeft:     dfs(rows[:rowMid], left, colMid),
					TopRight:    dfs(rows[:rowMid], colMid, right),
					BottomLeft:  dfs(rows[rowMid:], left, colMid),
					BottomRight: dfs(rows[rowMid:], colMid, right),
				}
			}
		}
	}

	// 叶子节点
	return &QuadTreeNode{
		Val:    rows[0][left] == 1,
		IsLeaf: true,
	}
}
