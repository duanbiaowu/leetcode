package leetcode

func repeatedCharacter(s string) byte {
	var res byte
	cntMap := make(map[byte]uint8)

	for i := range s {
		cntMap[s[i]]++
		if cntMap[s[i]] > 1 {
			return s[i]
		}
	}

	return res
}
