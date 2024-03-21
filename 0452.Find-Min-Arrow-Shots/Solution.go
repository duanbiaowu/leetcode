package leetcode

import "sort"

// 题目描述太过蛋疼！
// 本质上就是 **求解可以合并后的区间数量，也就是不重叠的区间数量**。
// 当然，本题中的区间合并的前提是：必须重叠才可以合并

// 写的更详细一些
// 计算出参数排序后的值
// 例如输入：points = [[10,16],[2,8],[1,6],[7,12]]
// 然后画出对应的 X 轴区间图，以右区间为准进行计算
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		// 以区间右边界作为排序基准
		return points[i][1] < points[j][1]
	})

	right := points[0][1]
	res := 1

	for _, p := range points {
		if p[0] > right {
			// 如果当前区间比之前的右边届区间大
			// 说明两个区间不存在重叠，需要多一支箭
			right = p[1]
			res++
		}
	}

	return res
}
