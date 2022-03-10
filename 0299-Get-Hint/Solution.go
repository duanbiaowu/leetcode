package leetcode

import (
	"strconv"
)

// secret.length == guess.length
// secret 和 guess 仅由数字组成
func getHint(secret string, guess string) string {
	nums := [10]int{}
	cntA, cntB := 0, 0
	for i := range secret {
		if secret[i] == guess[i] {
			cntA++
		} else {
			if nums[guess[i]-'0'] > 0 {
				cntB++
			}
			nums[guess[i]-'0']--
			if nums[secret[i]-'0'] < 0 {
				cntB++
			}
			nums[secret[i]-'0']++
		}
	}
	return strconv.Itoa(cntA) + "A" + strconv.Itoa(cntB) + "B"
}
