package leetcode

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 直接以第一个数组元素字符串为基准进行对比即可
	// 当然这里可以使用任意的数组元素作为基准字符串
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	// 代码执行到这里，说明第一个数组元素字符串本身就是最长前缀
	return strs[0]
}
