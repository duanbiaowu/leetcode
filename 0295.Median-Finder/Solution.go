package leetcode

import "container/heap"

type MinHeap []int

type MedianFinder struct {
	// 小顶堆: 存储输入数字较大的一半
	minHeap *MinHeap
	// 大顶堆: 存储输入数字较小的一半
	maxHeap *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(MinHeap),
		maxHeap: new(MinHeap),
	}
}

func (m *MedianFinder) AddNum(num int) {
	// 保证小顶堆的元素都大于大顶堆的元素
	if m.minHeap.Len() > 0 && num > m.minHeap.Top() {
		heap.Push(m.minHeap, num)
	} else {
		// 注意: 存储的是负数
		// 因为 Go 语言的 container/heap 包是小顶堆，所以存储相反数
		heap.Push(m.maxHeap, -num)
	}

	// 转换为负数符合堆的规则
	if m.minHeap.Len()-m.maxHeap.Len() == 2 {
		heap.Push(m.maxHeap, -(heap.Pop(m.minHeap)).(int))
	} else if m.maxHeap.Len()-m.minHeap.Len() == 2 {
		heap.Push(m.minHeap, -(heap.Pop(m.maxHeap)).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	// 注意堆顶元素负数的处理
	if m.minHeap.Len() > m.maxHeap.Len() {
		return float64((*m.minHeap).Top())
	} else if m.minHeap.Len() < m.maxHeap.Len() {
		return -float64((*m.maxHeap).Top())
	}

	return float64((*m.minHeap).Top()-(*m.maxHeap).Top()) / float64(2)
}

// MinHeap 实现 container/heap 的接口
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

func (h *MinHeap) Top() int {
	return (*h)[0]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
