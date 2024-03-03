package leetcode

type MinStack struct {
	// 每个栈元素是一个长度为 2 的数组
	// 同时保存当前入栈元素和当前栈内最小元素
	// [0]: 入栈值  [1]: 栈内最小值
	stack [][2]int
}

func Constructor() MinStack {
	return MinStack{
		stack: [][2]int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, [2]int{val, val})
	} else {
		this.stack = append(this.stack, [2]int{val, min(val, this.GetMin())})
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type MinStackOpt struct {
	stack []int // 存储入栈的所有元素和最小元素的差值
	min   int   // 存储栈内最小元素
}

func ConstructorOpt() MinStackOpt {
	return MinStackOpt{
		stack: []int{},
		min:   0,
	}
}

func (this *MinStackOpt) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
	}

	// 存入入栈元素和最小值之间的差值
	this.stack = append(this.stack, val-this.min)
	// 更新最小值
	if this.min > val {
		this.min = val
	}
}

func (this *MinStackOpt) Pop() {
	// 取出栈顶差值元素
	diff := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]

	// 如果差值小于 0, 说明出栈的元素就是最小值，最小值现在发生了新的变化
	// 新的最小值 = abs(栈顶差值) + 当前最小值
	if diff < 0 {
		this.min -= diff
	}
}

func (this *MinStackOpt) Top() int {
	// 取出栈顶差值元素
	diff := this.stack[len(this.stack)-1]
	// 如果差值小于 0, 说明当前栈顶元素就是最小值，直接返回
	if diff < 0 {
		return this.min
	}

	// 如果差值大于等于 0, 说明栈顶元素比当前最小值要大，加上差值返回即可
	return this.min + diff
}

func (this *MinStackOpt) GetMin() int {
	return this.min
}
