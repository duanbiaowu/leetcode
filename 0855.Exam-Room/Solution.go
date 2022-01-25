package leetcode

import "container/list"

type ExamRoom struct {
	// 坐着的同学的位
	seat *list.List
	n    int
}

// 一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1
func Constructor(n int) ExamRoom {
	return ExamRoom{list.New(), n - 1}
}

// 当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。
// 如果有多个这样的座位，他会坐在编号最小的座位上。
// (另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

// 核心是不断判断两点间距离，最终求得最大值
func (this *ExamRoom) Seat() int {
	// 还没有人入座，直接将 0 插入
	if this.seat.Len() == 0 {
		this.seat.PushFront(0)
		return 0
	}

	cur := this.seat.Front()
	pre := cur.Value.(int)
	// 头部特殊判断
	max := pre
	addVal := 0
	addFront := cur

	for cur = cur.Next(); cur != nil; cur = cur.Next() {
		val := cur.Value.(int)
		// 两点之间最远距离
		distant := (val - pre) / 2
		if distant > max {
			max = distant
			// 需要插入位置的后一个元素
			addFront = cur
			// 需要插入位置
			addVal = pre + distant
		}
		pre = val
	}

	// 尾部特殊判断
	distant := this.n - pre
	if distant > max {
		// 直接插入到链表尾部
		this.seat.PushBack(this.n)
		return this.n
	}

	this.seat.InsertBefore(addVal, addFront)
	return addVal
}

func (this *ExamRoom) Leave(p int) {
	for e := this.seat.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			this.seat.Remove(e)
			return
		}
	}
}
