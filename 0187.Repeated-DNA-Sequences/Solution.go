package leetcode

func findRepeatedDnaSequences(s string) []string {
	var res []string
	seen := map[string]int{}

	for i := 0; i <= len(s)-10; i++ {
		cur := s[i : i+10]
		if val, ok := seen[cur]; ok && val == 1 {
			res = append(res, cur)
		}
		seen[cur]++
	}

	return res
}
