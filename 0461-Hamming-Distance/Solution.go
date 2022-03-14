package leetcode

import "math/bits"

func hammingDistance(x int, y int) int {
	cnt := 0
	for s := x ^ y; s > 0; s &= s - 1 {
		cnt++
	}
	return cnt
}

func hammingDistance2(x int, y int) int {
	cnt := 0
	for s := x ^ y; s > 0; s >>= 1 {
		cnt += s & 1
	}
	return cnt
}

func hammingDistance3(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
