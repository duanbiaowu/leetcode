package leetcode

func isUnique(astr string) bool {
	mask := 0

	for i := range astr {
		offset := astr[i] - 'a'
		if mask&(1<<offset) > 0 {
			return false
		}

		mask |= 1 << offset
	}

	return true
}
