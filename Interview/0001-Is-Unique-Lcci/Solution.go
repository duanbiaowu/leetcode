package leetcode

func isUnique(astr string) bool {
	mark := 0
	for _, s := range astr {
		offset := s - 'a'
		if mark&(1<<offset) > 0 {
			return false
		}
		mark |= 1 << offset
	}
	return true
}
