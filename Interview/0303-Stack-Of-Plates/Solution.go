package leetcode

// 基础栈结构
type Stack struct {
	nums  []int
	point int
}

type StackOfPlates struct {
	// 除了最后一个栈，其余的都应该是满的
	stacks []*Stack
	// 每个栈容量上限
	cap int
	// 指向最后一个栈
	index int
}

func Constructor(cap int) StackOfPlates {
	return StackOfPlates{
		stacks: make([]*Stack, 0),
		cap:    cap,
		index:  0,
	}
}

func (this *StackOfPlates) Push(val int) {
	if this.cap == 0 {
		return
	}
	// 已有的栈全部满，扩容
	if this.index == len(this.stacks) {
		this.stacks = append(this.stacks, &Stack{
			nums:  make([]int, this.cap),
			point: 0,
		})
	}

	// Push 到最后一个栈
	stack := this.stacks[this.index]
	stack.nums[stack.point] = val
	stack.point++
	if stack.point == this.cap {
		this.index++
	}
}

func (this *StackOfPlates) Pop() int {
	if this.cap == 0 {
		return -1
	}
	i := this.index
	// 边界判断
	for ; i >= 0; i-- {
		if i == len(this.stacks) {
			continue
		}
		if this.stacks[i].point != 0 {
			break
		}
	}
	// 没有可用的栈
	if i < 0 {
		this.index = 0
		this.stacks = this.stacks[:0]
		return -1
	}

	stack := this.stacks[i]
	res := stack.nums[stack.point-1]
	stack.point--
	if stack.point == 0 {
		this.stacks = this.stacks[:i]
		if i > 0 && this.stacks[i-1].point < this.cap {
			this.index = i - 1
		} else {
			this.index = i
		}
	}
	return res
}

func (this *StackOfPlates) PopAt(index int) int {
	if this.cap == 0 {
		return -1
	}
	n := len(this.stacks)
	if index < 0 || index >= n {
		return -1
	}

	stack := this.stacks[index]
	res := stack.nums[stack.point-1]
	stack.point--
	if stack.point == 0 {
		if index < this.index {
			this.index--
			this.stacks = append(this.stacks[:index], this.stacks[index+1:]...)
		} else {
			this.stacks = this.stacks[:index]
		}
	}
	return res
}

/**
 * Your StackOfPlates object will be instantiated and called as such:
 * obj := Constructor(cap);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAt(index);
 */
