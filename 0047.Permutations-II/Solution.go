package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	visited := make([]bool, len(nums))

	// 去重前置排序
	sort.Ints(nums)

	backtrack(nums, &path, visited, &res)
	return res
}

func backtrack(nums []int, path *[]int, visited []bool, res *[][]int) {
	if len(*path) == len(nums) {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i, v := range nums {
		if visited[i] {
			continue
		}
		if i > 0 && !visited[i-1] && v == nums[i-1] {
			continue
		}

		visited[i] = true
		*path = append(*path, v)

		backtrack(nums, path, visited, res)

		visited[i] = false
		*path = (*path)[:len(*path)-1]
	}
}
