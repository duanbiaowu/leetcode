package leetcode

func twoSum(numbers []int, target int) []int {
	res := []int{}
	for left, right := 0, len(numbers)-1; left < right; {
		sum := numbers[left] + numbers[right]
		if sum == target {
			// 题目声明下标从 1 开始，所以这里将两个索引各加 1
			return []int{left + 1, right + 1}
		}
		if sum > target {
			right--
		} else {
			left++
		}
	}
	return res
}
