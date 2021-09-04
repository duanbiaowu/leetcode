package leetcode

type MinStack struct {
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
