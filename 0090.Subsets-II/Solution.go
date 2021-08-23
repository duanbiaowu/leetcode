package leetcode

import (
	"sort"
)

// a = {5, 2, 9}
// 0/1 序列对应的二进制数正好从 0 到 2^n - 1
// ---0/1序列------------子集--------------0/1序列对应的二进制数------------
/**
      000               {}                0
 	  001				{9}				  1
	  010				{2}				  2
      011				{2, 9}			  3
      100				{5}				  4
      101				{5, 9}			  5
      110				{5, 2}			  6
      111				{5, 2, 9}		  7
*/
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		var set []int
		flag := true
		for i, v := range nums {
			if mask>>i&1 > 0 {
				// 如果当前元素和上一个元素相同 && 上一个元素未被选择
				if i > 0 && v == nums[i-1] && mask>>(i-1)&1 == 0 {
					flag = false
					break
				}
				set = append(set, v)
			}
		}

		if flag {
			res = append(res, append([]int{}, set...))
		}
	}
	return res
}

func subsetsRecursively(nums []int) [][]int {
	var path []int
	var res [][]int
	sort.Ints(nums)
	backtrack(nums, 0, &path, &res)
	return res
}

func backtrack(nums []int, begin int, path *[]int, res *[][]int) {
	*res = append(*res, append([]int{}, *path...))
	for i := begin; i < len(nums); i++ {
		if i > begin && nums[i] == nums[i-1] {
			continue
		}
		*path = append(*path, nums[i])
		backtrack(nums, i+1, path, res)
		*path = (*path)[:len(*path)-1]
	}
}
