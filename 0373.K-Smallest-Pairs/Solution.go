package leetcode

import (
	"container/heap"
	"sort"
)

// 超出内存限制
func kSmallestPairsSample(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return nil
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			res = append(res, []int{nums1[i], nums2[j]})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})

	if k > len(res) {
		return nil
	}
	return res[:k]
}

// 超出时间限制
func kSmallestPairsSample2(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return nil
	}

	res := make([][]int, 0)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			res = append(res, []int{nums1[i], nums2[j]})

			if len(res) > k {
				sort.Slice(res, func(i, j int) bool {
					return res[i][0]+res[i][1] < res[j][0]+res[j][1]
				})
				res = res[:k]
			}
		}
	}

	return res
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	h := hp{nums1: nums1, nums2: nums2}

	// 题目要求: 其中第一个元素来自 nums1，第二个元素来自 nums2
	// 并且两个参数数组已经是有序的
	// 所以这里只需要将 nums1 的前 k 个元素加入到堆中即可
	// 因为结果集合必然由 nums1 的前 k 个元素与 nums2 的前 k 个元素组合形成
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}

	ans := make([][]int, 0)
	for h.Len() > 0 && len(ans) < k {
		// 最小堆, 每次取出堆顶元素即为最小值
		top := heap.Pop(&h).(pair)
		i, j := top.i, top.j

		// 将最小值加入到结果集中
		ans = append(ans, []int{nums1[i], nums2[j]})

		// 如果 j+1 < n, 则将 nums1[i] 与 nums2[j+1] 加入到堆中
		// 形成新的 “预选最小值” 组合
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}

	return ans
}

type pair struct {
	i, j int
}

type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h *hp) Len() int {
	return len(h.data)
}

func (h *hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(pair))
}

func (h *hp) Pop() interface{} {
	top := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return top
}

func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
