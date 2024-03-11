package leetcode

import (
	"sort"
)

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := hammingWeight(uint32(arr[i])), hammingWeight(uint32(arr[j]))
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}

func sortByBitsOpts(arr []int) []int {
	// 递推预处理
	// 定义 bitMemo[i] 为数字 i 的二进制中为 1 的数量
	// 则可以列出下面的递推式
	// bitMemo[i] = bitMemo[i>>1] + i&1
	bitMemo := [1e4 + 1]int{}
	for i := 1; i <= 1e4; i++ {
		bitMemo[i] = bitMemo[i>>1] + i&1
	}

	sort.Slice(arr, func(i, j int) bool {
		x, y := bitMemo[arr[i]], bitMemo[arr[j]]
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}
