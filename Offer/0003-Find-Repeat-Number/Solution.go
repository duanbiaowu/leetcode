package leetcode

// 在一个长度为 n 的数组 nums 里的所有数字都在 0 ~ n-1 的范围内
// 此说明含义：数组元素的 索引 和 值 是 一对多 的关系（数组中某些数字是重复的）
// 因此，可遍历数组并通过交换操作
// 使元素的 索引 与 值一一对应（ 即 nums[i] = i ）
// 通过索引映射对应的值，起到与字典等价的作用

// 遍历中，第一次遇到数字 xx 时，将其交换至索引 x 处
// 而当第二次遇到数字 x 时，一定有 nums[x] = x ，此时即可得到一组重复数字
func findRepeatNumber(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i++
		} else {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
