package leetcode

func majorityElement(nums []int) int {
	// “多数元素” 出现次数通过计数器表示
	// 不论过程如何变化，多数元素的出现规律都是 “此消彼长”
	// “多数元素” 使计数器增加
	// “其他元素“ 使计数器减少
	// 因为 “多数元素” 出现次数大于数组长度/2
	// 所以最终计数器的值一定是总数
	target, cnt := 0, 0

	for _, n := range nums {
		// 计数器归零时，设置目标值为当前元素
		if cnt == 0 {
			target = n
		}
		// 如果当前元素和目标值相同，计数器+1
		// 否则 计数器-1
		if target == n {
			cnt++
		} else {
			cnt--
		}
	}

	// 返回目标值
	return target
}
