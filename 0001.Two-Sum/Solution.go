package leetcode

// 暴力解法
//func TwoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i] + nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return []int{}
//}

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		if diff, ok := m[target-num]; ok {
			return []int{diff, i}
		}
		m[num] = i
	}

	return []int{}
}
