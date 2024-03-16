package leetcode

import "sort"

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
