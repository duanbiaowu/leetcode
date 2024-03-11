package leetcode

func findErrorNums(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	// sum(nums) - sum(set(nums)) = 重复的数字
	// (1 + len(nums)) * len(nums) / 2 - sum(set(nums)) = 丢失的数字

	set := make(map[int]struct{})
	sum, setSum := 0, 0
	for _, num := range nums {
		if _, ok := set[num]; !ok {
			setSum += num
		}
		set[num] = struct{}{}
		sum += num
	}

	n := len(nums)
	return []int{sum - setSum, (n*(n+1))>>1 - setSum}
}
