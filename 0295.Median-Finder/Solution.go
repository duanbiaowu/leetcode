package leetcode

// 自定义 Comparator 函数类型
// 用于控制 (大/小) 堆的排序方式
type Comparator func(x, y int) bool

type Heap struct {
	data       []int
	comparator Comparator
}

// 依赖注入对象创建新的堆
func NewHeap(comparator Comparator) *Heap {
	return &Heap{
		data:       []int{},
		comparator: comparator,
	}
}

// 向堆中插入元素
func (h *Heap) push(val int) {
	h.data = append(h.data, val)

	// 从最后一个节点开始 (上浮) 操作
	i := len(h.data) - 1
	for i > 0 {
		parent := (i - 1) / 2
		if !h.comparator(h.data[i], h.data[parent]) {
			break
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

// 弹出堆顶元素 (下沉) 操作
func (h *Heap) pop() int {
	if len(h.data) == 0 {
		return -1
	}

	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.heapifyDown(0)

	return top
}

// 堆的调整操作
func (h *Heap) heapifyDown(i int) {
	n := len(h.data)
	pivot := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && h.comparator(h.data[left], h.data[pivot]) {
		pivot = left
	}
	if right < n && h.comparator(h.data[right], h.data[pivot]) {
		pivot = right
	}

	if pivot != i {
		h.data[i], h.data[pivot] = h.data[pivot], h.data[i]
		h.heapifyDown(pivot)
	}
}

// 获取堆顶元素
func (h *Heap) top() int {
	if len(h.data) == 0 {
		return -1
	}
	return h.data[0]
}

// MedianFinder 数据流中位数结构
type MedianFinder struct {
	lower *Heap // 大顶堆: 存储输入数字较小的一半
	upper *Heap // 小顶堆: 存储输入数字较大的一半
}

// 构造函数
func Constructor() MedianFinder {
	return MedianFinder{
		lower: NewHeap(func(a, b int) bool { return a > b }), // 大顶堆
		upper: NewHeap(func(a, b int) bool { return a < b }), // 小顶堆
	}
}

// 添加一个数字到堆中
func (mf *MedianFinder) AddNum(num int) {
	// 优先将数字插入到小顶堆
	// 如果小顶堆为空，或者插入的数字小于 小顶堆 的对顶元素
	// 将数字插入到小顶堆
	if len(mf.lower.data) == 0 || num <= mf.lower.top() {
		mf.lower.push(num)
	} else {
		mf.upper.push(num)
	}

	// 平衡两个堆
	// 如果两个堆的元素数量相同，不做任何处理
	// 否则永远保持 小顶堆的元素数量大于等于 大顶堆
	//    便于计算中位数，元素数量为奇数时直接返回小顶堆根元素即可
	if len(mf.lower.data)-len(mf.upper.data) > 1 {
		mf.upper.push(mf.lower.pop())
	} else if len(mf.upper.data) > len(mf.lower.data) {
		mf.lower.push(mf.upper.pop())
	}
}

// 获取当前中位数
func (mf *MedianFinder) FindMedian() float64 {
	// 两个堆的元素数量只会存在 2 种情况
	// 1. 小顶堆数量 > 大顶堆数量
	// 2. 小顶堆数量 = 大顶堆数量
	if len(mf.lower.data) > len(mf.upper.data) {
		return float64(mf.lower.top())
	}

	return float64(mf.lower.top()+mf.upper.top()) / 2.0
}
