package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	var visited = make([]bool, len(nums))
	// 去重前置排序
	sort.Ints(nums)
	backtrack(nums, 0, &path, &visited, &res)
	return res
}

func backtrack(nums []int, begin int, path *[]int, visited *[]bool, res *[][]int) {
	if begin == len(nums) {
		*res = append(*res, append([]int{}, *path...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*visited)[i] {
			// 去重
			if i > 0 && nums[i] == nums[i-1] && !(*visited)[i-1] {
				continue
			}
			(*visited)[i] = true
			*path = append(*path, nums[i])
			backtrack(nums, begin+1, path, visited, res)
			(*visited)[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}
