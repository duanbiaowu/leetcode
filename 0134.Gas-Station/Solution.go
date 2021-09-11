package leetcode

// https://leetcode-cn.com/problems/gas-station/comments/21285
// https://leetcode-cn.com/problems/gas-station/solution/shi-yong-tu-de-si-xiang-fen-xi-gai-wen-ti-by-cyayc/
func canCompleteCircuit(gas []int, cost []int) int {
	rest, run, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		run += gas[i] - cost[i]
		rest += gas[i] - cost[i]
		if run < 0 {
			start = i + 1
			run = 0
		}
	}
	if rest < 0 {
		return -1
	}
	return start
}
