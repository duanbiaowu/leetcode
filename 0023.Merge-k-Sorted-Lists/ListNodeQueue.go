package leetcode

type ListNodeQueue []*ListNode

func (queue ListNodeQueue) Len() int {
	return len(queue)
}

func (queue ListNodeQueue) Less(i, j int) bool {
	if queue[i] == nil {
		return false
	}
	if queue[j] == nil {
		return true
	}
	return queue[i].Val < queue[j].Val
}

func (queue ListNodeQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
}

func (queue *ListNodeQueue) Push(x interface{}) {
	*queue = append(*queue, x.(*ListNode))
}

func (queue *ListNodeQueue) Pop() interface{} {
	n := len(*queue)
	top := (*queue)[n-1]
	*queue = (*queue)[0 : n-1]
	return top
}
