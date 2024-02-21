package leetcode

func findTheDifference(s string, t string) byte {
	sMap := make(map[int32]int)

	// 细节: 循环时，v 的类型是 int32
	for _, v := range s {
		sMap[v]++
	}

	for _, v := range t {
		if _, ok := sMap[v]; !ok {
			return byte(v)
		}

		sMap[v]--
		if sMap[v] < 0 {
			return byte(v)
		}
	}

	return ' '
}

func findTheDifferenceOpt(s string, t string) byte {
	var sum uint8

	// 细节: 循环时，s[i] 的类型是 uint8
	for i := range s {
		sum += s[i]
	}
	for i := range t {
		sum -= t[i]
	}

	return -sum
}

func findTheDifferenceOpt2(s string, t string) byte {
	var res byte
	if len(t) == 0 {
		return res
	}

	for i := range s {
		res ^= s[i] ^ t[i]
	}

	return res ^ t[len(t)-1]
}
