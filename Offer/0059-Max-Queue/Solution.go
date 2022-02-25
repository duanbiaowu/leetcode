package leetcode

type MaxQueue struct {
	queue []int
	max   []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

// 队列：FIFO
// 比如入队为 3, 4, 5, 6，那么最大值为 6
// 6 未出队之前，Max_Value 保持不变，前面的值可以直接舍弃
// 当然还有一种办法就是将前面的值全部置为 6，但是会多浪费一些空间
func (this *MaxQueue) Push_back(value int) {
	this.queue = append(this.queue, value)
	for len(this.max) > 0 && this.max[len(this.max)-1] <= value {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	n := this.Top()
	this.queue = this.queue[1:]
	if this.Max_value() == n {
		this.max = this.max[1:]
	}
	return n
}

func (this *MaxQueue) Top() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}
