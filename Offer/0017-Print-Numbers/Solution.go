package leetcode

var NUMBERS = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func printNumbers(n int) []int {
	var res, path []int
	for i := 1; i <= n; i++ {
		backtrack(0, i, &path, &res)
	}
	return res
}

func backtrack(begin, n int, path *[]int, res *[]int) {
	if begin == n {
		*res = append(*res, convertSliceToInt(*path))
		return
	}
	start := 0
	if begin == 0 {
		start = 1
	}
	for i := start; i < 10; i++ {
		*path = append(*path, NUMBERS[i])
		backtrack(begin+1, n, path, res)
		*path = (*path)[:len(*path)-1]
	}
}

func convertSliceToInt(nums []int) int {
	n := len(nums)
	if n < 1 {
		return 0
	}
	res := nums[0]
	for i := 1; i < n; i++ {
		res = res*10 + nums[i]
	}
	return res
}
