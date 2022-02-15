package leetcode

import "container/heap"

// 利用一个大顶堆一个小顶堆，中位数可以看作是隔开两个数组的分位线
// 左端用大顶堆，右端用小顶堆，最后两个堆（奇数个是一个）的peek求平均即可
// 大顶堆逆序排列，小顶堆正序排列
type MedianFinder struct {
	max MaxHeap
	min MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	min := MinHeap{}
	max := MaxHeap{}
	heap.Init(&min)
	heap.Init(&max)
	return MedianFinder{max, min}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.max, num)
	heap.Push(&this.min, heap.Pop(&this.max))

	if this.max.Len() < this.min.Len() {
		heap.Push(&this.max, heap.Pop(&this.min))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.max.Len() > this.min.Len() {
		return float64(this.max.Top().(int))
	}
	return float64(this.max.Top().(int)+this.min.Top().(int)) / 2
}

type Heap []int

func (this *Heap) Top() interface{} {
	return (*this)[0]
}

func (this Heap) Len() int {
	return len(this)
}

func (this Heap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *Heap) Push(i interface{}) {
	*this = append(*this, i.(int))
}

func (this *Heap) Pop() interface{} {
	top := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return top
}

type MaxHeap struct {
	Heap
}

type MinHeap struct {
	Heap
}

func (this MaxHeap) Less(i, j int) bool {
	return this.Heap[i] > this.Heap[j]
}

func (this MinHeap) Less(i, j int) bool {
	return this.Heap[i] < this.Heap[j]
}
