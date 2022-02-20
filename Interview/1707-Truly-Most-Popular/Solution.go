package leetcode

import (
	"strconv"
	"strings"
)

type Union struct {
	parent []int
	nums   []int
	names  []string
}

func Constructor(names []string) *Union {
	parent := make([]int, len(names))
	nums := make([]int, len(names))

	for i := range names {
		low := strings.Index(names[i], "(")
		hi := strings.Index(names[i], ")")
		v, _ := strconv.Atoi(names[i][low+1 : hi])
		nums[i] = v
		parent[i] = i
	}

	return &Union{parent, nums, names}
}

func (u *Union) find(x int) int {
	for u.parent[x] != x {
		// 路径压缩
		u.parent[x] = u.parent[u.parent[x]]
		x = u.parent[x]
	}
	return x
}

func (u *Union) join(x, y int) {
	rootP := u.find(x)
	rootQ := u.find(y)
	if rootP == rootQ {
		return
	}
	if u.names[rootP] < u.names[rootQ] {
		u.parent[rootQ] = rootP
		u.nums[rootP] += u.nums[rootQ]
	} else {
		u.parent[rootP] = rootQ
		u.nums[rootQ] += u.nums[rootP]
	}
}

func trulyMostPopular(names []string, synonyms []string) []string {
	var res []string
	m := make(map[string]int)

	for i := range names {
		n := strings.Index(names[i], "(")
		m[names[i][:n]] = i
	}

	union := Constructor(names)

	for i := range synonyms {
		low := strings.Index(synonyms[i], "(")
		mid := strings.Index(synonyms[i], ",")
		hi := strings.Index(synonyms[i], ")")
		s1 := synonyms[i][low+1 : mid]
		s2 := synonyms[i][mid+1 : hi]
		union.join(m[s1], m[s2])
	}

	var buf strings.Builder
	for i := range union.parent {
		if union.parent[i] == i {
			n := strings.Index(names[i], "(")
			buf.Reset()
			buf.WriteString(names[i][:n])
			buf.WriteString("(")
			buf.WriteString(strconv.Itoa(union.nums[i]))
			buf.WriteString(")")
			res = append(res, buf.String())
		}
	}

	return res
}
