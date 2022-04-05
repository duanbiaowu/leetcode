package leetcode

import (
	"container/heap"
)

type element struct {
	val int
	cnt int
}

type priorityQueue []*element

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

// Less 小顶堆
func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i].cnt < (*pq)[j].cnt
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*element))
}

// Pop 小顶堆
func (pq *priorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}

	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}

	q := &priorityQueue{}
	heap.Init(q)
	for val, cnt := range m {
		heap.Push(q, &element{val, cnt})
		if (*q).Len() > k {
			heap.Pop(q)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(q).(*element).val
	}
	return res
}

// 桶排序
func topKFrequent2(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}

	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]]++
	}

	cntMap := make(map[int][]int)
	for val, cnt := range numMap {
		if _, ok := cntMap[cnt]; ok {
			cntMap[cnt] = append(cntMap[cnt], val)
		} else {
			cntMap[cnt] = []int{val}
		}
	}

	res := make([]int, k)
	pos := 0
	for i := len(nums); i > 0; i-- {
		if _, ok := cntMap[i]; ok {
			for j := range cntMap[i] {
				res[pos] = cntMap[i][j]
				pos++
				if pos == k {
					return res
				}
			}
		}
	}
	return res
}
