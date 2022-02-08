package leetcode

const (
	N = 3
)

type TripleInOne struct {
	data    []int
	indexes []int
	size    int
}

func Constructor(stackSize int) TripleInOne {
	indexes := make([]int, N)
	for i := 0; i < N; i++ {
		indexes[i] = i * stackSize
	}
	return TripleInOne{
		data:    make([]int, N*stackSize),
		indexes: indexes,
		size:    stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	idx := this.indexes[stackNum]
	if idx < (stackNum+1)*this.size {
		this.data[idx] = value
		this.indexes[stackNum]++
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	idx := this.indexes[stackNum]
	if idx > stackNum*this.size {
		this.indexes[stackNum]--
		return this.data[idx-1]
	}
	return -1
}

func (this *TripleInOne) Peek(stackNum int) int {
	idx := this.indexes[stackNum]
	if idx > stackNum*this.size {
		return this.data[idx-1]
	}
	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.indexes[stackNum] == stackNum*this.size
}
