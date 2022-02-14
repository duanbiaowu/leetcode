package leetcode

func maxAliveYear(birth []int, death []int) int {
	if len(birth) == 0 || len(death) == 0 {
		return 0
	}
	// 假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）
	cnt := [102]int{}
	// 出生年份 +1
	// 死亡年份下一年 -1（点睛之笔）
	for i := 0; i < len(birth); i++ {
		cnt[birth[i]-1900]++
		cnt[death[i]-1900+1]--
	}
	// 前缀和即当前人数, 求最大值
	tmp, max, index := 0, 0, 0
	// 2000年死的，不计入统计
	for i := 0; i < 101; i++ {
		tmp += cnt[i]
		if tmp > max {
			max = tmp
			index = i
		}
	}
	return index + 1900
}
