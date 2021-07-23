package leetcode

var roman = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	res, n := 0, len(s)
	for i := 0; i < n; i++ {
		val := roman[s[i]]
		// 介于两个数字之间: IV => 4, IX => 9
		if i < n-1 && val < roman[s[i+1]] {
			res -= val
		} else {
			res += val
		}
	}
	return res
}
