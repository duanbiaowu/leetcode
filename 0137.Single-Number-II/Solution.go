package leetcode

func singleNumber(nums []int) int {
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		a = (a ^ nums[i]) & ^b
		b = (b ^ nums[i]) & ^a
	}
	return a
}
