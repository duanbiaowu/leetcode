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
	for i, v := range nums {
		if index, ok := m[target - v]; ok {
			return []int{index, i}
		}
		m[v] = i
	}
	return []int{}
}