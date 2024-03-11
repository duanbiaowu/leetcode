package leetcode

func masterMind(solution string, guess string) []int {
	hit, total := 0, 0
	if len(solution) == 0 || len(guess) == 0 {
		return []int{hit, total}
	}

	m := [26]int{}
	for i := 0; i < 4; i++ {
		m[guess[i]-'A']++
	}
	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			hit++
		}
		if m[solution[i]-'A'] > 0 {
			total++
			m[solution[i]-'A']--
		}
	}
	// 只猜对颜色但槽位猜错，算一次“伪猜中”
	// 总次数 - 总猜中 = 总伪猜中
	return []int{hit, total - hit}
}
