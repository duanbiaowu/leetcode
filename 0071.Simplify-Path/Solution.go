package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack []string

	for _, p := range strings.Split(path, "/") {
		// 无效路径字符，直接跳过
		if p == "" || p == "." {
			continue
		}

		// 目录名称/文件名称 入栈
		if p != ".." {
			stack = append(stack, p)
		} else if p == ".." && len(stack) > 0 {
			// 如果是上级目录 ..
			// 将栈顶元素出栈
			stack = stack[:len(stack)-1]
		}
	}

	return "/" + strings.Join(stack, "/")
}
