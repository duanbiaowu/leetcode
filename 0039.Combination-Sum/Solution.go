package leetcode

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	var path []int
	var res [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, 0, path, &res)
	return res
}

func backtrack(nums []int, target, begin int, path []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int{}, path...))
		return
	}

	for i := begin; i < len(nums) && nums[i] <= target; i++ {
		path = append(path, nums[i])
		backtrack(nums, target-nums[i], i, path, res)
		path = path[:len(path)-1]
	}
}
