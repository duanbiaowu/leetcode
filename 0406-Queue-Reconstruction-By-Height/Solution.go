package leetcode

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	// 身高顺序，
	// 如果身高相同，人数倒序
	sort.Slice(people, func(i, j int) bool {
		x, y := people[i], people[j]
		return x[0] < y[0] || x[0] == y[0] && x[1] > y[1]
	})

	res := make([][]int, len(people))
	for i := range people {
		pos := people[i][1] + 1
		for j := range res {
			if res[j] == nil {
				pos--
				if pos == 0 {
					res[j] = people[i]
					break
				}
			}
		}
	}
	return res
}
