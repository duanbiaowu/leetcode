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

	// 统计每个数字出现的次数
	// key: 数字, value: 出现的次数
	// 例如: nums = [1, 1, 1, 2, 2, 3, 4, 5],
	// numMap = {1: 3, 2: 2, 3: 1, 4: 1, 5: 1}
	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]]++
	}

	// 按照数字出现的次数进行桶排序
	// key: 出现的次数, value: 出现次数为 key 的数字列表
	// 例如: cntMap = {3: [1], 2: [2], 1: [3, 4, 5]}
	cntMap := make(map[int][]int)
	for num, cnt := range numMap {
		cntMap[cnt] = append(cntMap[cnt], num)
	}

	res := []int{}

	// 题目要求: 1 <= nums.length <= 105
	// 从后往前遍历, 保证出现次数多的数字在前面
	for i := len(nums); i > 0; i-- {
		if _, ok := cntMap[i]; !ok {
			continue
		}

		for _, num := range cntMap[i] {
			res = append(res, num)
			// 题目要求返回前 k 个出现次数最多的数字
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
