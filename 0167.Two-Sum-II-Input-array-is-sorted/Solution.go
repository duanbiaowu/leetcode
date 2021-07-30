package leetcode

func twoSum(numbers []int, target int) []int {
	res := []int{}
	for low, hi := 0, len(numbers)-1; low < hi; {
		sum := numbers[low] + numbers[hi]
		if sum == target {
			return []int{low + 1, hi + 1}
		}
		if sum > target {
			hi--
		} else {
			low++
		}
	}
	return res
}
