package leetcode

import "bytes"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	// 按照与逐行读取 Z 字形图案相同的顺序访问字符串
	var buf bytes.Buffer
	n := len(s)

	// 行 0 中的字符位于索引 k:(2*numRows−2)
	// 行 numRows−1 中的字符位于索引 k:(2⋅numRows−2)+numRows−1
	// 内部的 行 i 中的字符位于索引   k:(2⋅numRows−2)+i 以及 (k+1)(2⋅numRows−2)−i
	cycleLen := 2 * numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j + i < n; j += cycleLen {
			buf.WriteByte(s[j+i])
			if i != 0 && i != numRows - 1 && j + cycleLen - i < n {
				buf.WriteByte(s[j + cycleLen - i])
			}
		}
	}
	return buf.String()
}