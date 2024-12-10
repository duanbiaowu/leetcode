package leetcode

import (
	"strings"
)

// 按行访问
//func convert(s string, numRows int) string {
//	if numRows <= 1 {
//		return s
//	}
//
//	// 按照与逐行读取 Z 字形图案相同的顺序访问字符串
//	var buf bytes.Buffer
//	n := len(s)
//
//	// 行 0 中的字符位于索引 k:(2*numRows−2)
//	// 行 numRows−1 中的字符位于索引 k:(2⋅numRows−2)+numRows−1
//	// 内部的 行 i 中的字符位于索引   k:(2⋅numRows−2)+i 以及 (k+1)(2⋅numRows−2)−i
//	cycleLen := 2 * numRows - 2
//
//	for i := 0; i < numRows; i++ {
//		for j := 0; j + i < n; j += cycleLen {
//			buf.WriteByte(s[j+i])
//			if i != 0 && i != numRows - 1 && j + cycleLen - i < n {
//				buf.WriteByte(s[j + cycleLen - i])
//			}
//		}
//	}
//	return buf.String()
//}

// 按行排序
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	// 注意结构: 因为返回结果为字符串, 所以不需要矩阵存储
	rows := []strings.Builder{}
	for i := 0; i < min(numRows, len(s)); i++ {
		rows = append(rows, strings.Builder{})
	}

	curRow := 0
	goingDown := false

	for i := range s {
		rows[curRow].WriteByte(s[i])
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	var sb strings.Builder
	for _, row := range rows {
		sb.WriteString(row.String())
	}
	return sb.String()
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
