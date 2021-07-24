package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	var res = [][]int{}
	if n < 4 {
		return res
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		// 第一个位置元素去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 前4个元素大于目标值, 说明没有满足条件的元素了
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		// 第一个位置元素增加
		if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// 第二个位置元素去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 前4个元素大于目标值, 说明没有满足条件的元素了
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			// 第二个位置元素增加
			if nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}

			for left, right := j+1, n-1; left < right; {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 第三个位置元素: 去重
					for left++; left < right && nums[left] == nums[left-1]; left++ {

					}
					// 第四个位置元素: 去重
					for right--; left < right && nums[right] == nums[right+1]; right-- {

					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return res
}
