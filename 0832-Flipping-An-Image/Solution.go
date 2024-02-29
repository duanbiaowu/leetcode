package leetcode

// 输入：														输出:
//
//	1,1,0					        0,1,1  						1,0,0
//	1,0,1    反转每一行 (逆序) ==>   1,0,1  反转 (0 <-> 1)  ==>   0,1,0
//	0,0,0							0,0,0					    1,1,1

// 因为数组中的元素只有 0 和 1
// 在进行逆序 + 反转两个操作后，元素就会恢复原样
// 需要注意是两个边界情况:
//  1. 数组的第 i 个元素和第 N - i 个元素相同
//  2. 数组长度奇数，需要单独处理中间的元素
func flipAndInvertImage(image [][]int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	for _, row := range image {
		left, right := 0, len(row)-1
		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}
			left++
			right--
		}

		if left == right {
			row[left] ^= 1
		}
	}

	return image
}
