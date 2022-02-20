package leetcode

import "strings"

var (
	singles   = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	teens     = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens      = []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands = []string{"", "Thousand", "Million", "Billion"}
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var sb strings.Builder
	for i, unit := 3, int(1e9); i >= 0; i-- {
		if curNum := num / unit; curNum > 0 {
			num -= curNum * unit
			sb.WriteString(toEnglish(curNum))
			sb.WriteString(thousands[i])
			sb.WriteByte(' ')
		}
		unit /= 1000
	}
	return strings.TrimSpace(sb.String())
}

func toEnglish(num int) string {
	var sb strings.Builder
	if num >= 100 {
		sb.WriteString(singles[num/100])
		sb.WriteString(" Hundred ")
		num %= 100
	}
	if num >= 20 {
		sb.WriteString(tens[num/10])
		sb.WriteByte(' ')
		num %= 10
	}
	if num > 0 && num < 10 {
		sb.WriteString(singles[num])
		sb.WriteByte(' ')
	} else if num >= 10 {
		sb.WriteString(teens[num-10])
		sb.WriteByte(' ')
	}
	return sb.String()
}
