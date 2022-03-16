package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	res, cnt := 0, 0
	left, right := 0, 0

	for right = range s {
		cnt += abs(int(s[right]-'0') - int(t[right]-'0'))
		for cnt > maxCost {
			cnt -= abs(int(s[left]-'0') - int(t[left]-'0'))
			left++
		}
		res = max(res, right-left+1)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
