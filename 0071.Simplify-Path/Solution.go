package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	for _, component := range strings.Split(path, "/") {
		if component != "" && component != "." && component != ".." {
			stack = append(stack, component)
		} else if component == ".." && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return "/" + strings.Join(stack, "/")
}
