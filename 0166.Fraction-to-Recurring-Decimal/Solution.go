package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func fractionToDecimal(numerator int, denominator int) string {
	// 分母为0
	if numerator == 0 {
		return "0"
	}
	// 分子为0
	if denominator == 0 {
		return ""
	}

	var res strings.Builder

	// (numerator > 0) ^ (denominator > 0):
	// Invalid operation: (numerator > 0) ^ (denominator > 0) (the operator ^ is not defined on bool)
	// 一正一负=负号
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		res.WriteByte('-')
	}

	// 分子分母全部转为正数
	numerator = abs(numerator)
	denominator = abs(denominator)

	// 整数部分
	res.WriteString(strconv.Itoa(numerator / denominator))

	numerator %= denominator
	// 余数为0表示整除，直接返回
	if numerator == 0 {
		return res.String()
	}

	// 小数部分
	res.WriteByte('.')

	// 小数点的下标, 默认没有循环余数
	index := -1
	// 记录循环余数开始下标
	record := map[int]int{}
	for numerator > 0 {
		// 余数扩大10倍, 开始求商, (草稿纸算法^_^)
		numerator *= 10
		// 出现循环余数
		if position, ok := record[numerator]; ok {
			index = position
			break
		} else {
			record[numerator] = res.Len()
		}

		res.WriteString(strconv.Itoa(numerator / denominator))
		numerator %= denominator
	}

	// 余数没有循环
	if index == -1 {
		return res.String()
	}

	// 在循环开始和结束加()
	s := res.String()
	return fmt.Sprintf("%s(%s)", s[:index], s[index:])
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
