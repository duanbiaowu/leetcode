package leetcode

import (
	"sort"
)

// reference: https://leetcode-cn.com/problems/circus-tower-lcci/solution/xiang-jie-tan-xin-er-fen-sou-suo-by-zrd-4/
func bestSeqAtIndex(height []int, weight []int) int {
	n := len(height)
	person := make([][2]int, n)
	for i := range height {
		person[i] = [2]int{height[i], weight[i]}
	}

	// 身高降序，若身高相同，体重升序
	sort.Slice(person, func(i, j int) bool {
		if person[i][0] == person[j][0] {
			return person[i][1] < person[j][1]
		}
		return person[i][0] > person[j][0]
	})

	// 可能形成最长子序列的元素
	res := 0
	for index := range person {
		// 在结果中找到第一个不能叠在 p 上面的人
		k := sort.Search(res, func(i int) bool {
			return person[i][0] <= person[index][0] || person[i][1] <= person[index][1]
		})

		// 将这个人替换成 p（节省空间）
		person[k] = person[index]
		if k == res {
			res++
		}
	}
	return res
}
