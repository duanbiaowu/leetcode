package leetcode

// reference: https://leetcode-cn.com/problems/majority-element-ii/solution/qiu-zhong-shu-ii-by-leetcode-solution-y1rn/
func majorityElement(nums []int) []int {
	elem1, elem2 := 0, 0
	vote1, vote2 := 0, 0

	for i := range nums {
		if vote1 > 0 && nums[i] == elem1 { // 如果该元素为第一个元素，则计数加1
			vote1++
		} else if vote2 > 0 && nums[i] == elem2 { // 如果该元素为第二个元素，则计数加1
			vote2++
		} else if vote1 == 0 { // 选择第一个元素
			elem1 = nums[i]
			vote1++
		} else if vote2 == 0 { // 选择第二个元素
			elem2 = nums[i]
			vote2++
		} else { // 如果三个元素均不相等，相互抵消一次
			vote1--
			vote2--
		}
	}

	// 统计出现总次数
	cnt1, cnt2 := 0, 0
	for i := range nums {
		if vote1 > 0 && nums[i] == elem1 {
			cnt1++
		}
		if vote2 > 0 && nums[i] == elem2 {
			cnt2++
		}
	}

	var res []int
	if cnt1 > 0 && cnt1 > len(nums)/3 {
		res = append(res, elem1)
	}
	if cnt2 > 0 && cnt2 > len(nums)/3 {
		res = append(res, elem2)
	}
	return res
}
