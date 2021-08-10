package leetcode

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	var visited = make([]bool, len(nums))
	backtrack(nums, &path, &visited, &res)
	return res
}

func backtrack(nums []int, path *[]int, visited *[]bool, res *[][]int) {
	if len(*path) == len(nums) {
		row := make([]int, len(nums))
		copy(row, *path)
		*res = append(*res, row)
		return
	}

	for i, v := range nums {
		if !(*visited)[i] {
			(*visited)[i] = true
			*path = append(*path, v)
			backtrack(nums, path, visited, res)
			(*visited)[i] = false
			*path = (*path)[:len(*path)-1]
		}
	}
}
