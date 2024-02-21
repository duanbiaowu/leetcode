package leetcode

import (
	"strconv"
)

func readBinaryWatch(turnedOn int) []string {
	res := make([]string, 0)

	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if hammingWeight(i)+hammingWeight(j) == turnedOn {
				val := strconv.Itoa(i) + ":"
				if j < 10 {
					val += "0"
				}
				val += strconv.Itoa(j)

				res = append(res, val)
			}
		}
	}

	return res
}

func hammingWeight(num int) int {
	cnt := 0
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return cnt
}
