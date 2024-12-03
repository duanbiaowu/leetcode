package leetcode

import "math/rand"

// 题目要求: 每个函数的平均时间复杂度为 O(1)
// 数据结构中必然存在 HashMap
// 除此之外，还要存储每次 Insert 的元素，额外使用一个辅助数组即可
type RandomizedSet struct {
	m    map[int]int
	data []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:    map[int]int{},
		data: []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; !ok {
		this.m[val] = len(this.data)
		this.data = append(this.data, val)

		return true
	}

	return false
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		tail, index := this.data[len(this.data)-1], this.m[val]
		this.m[tail], this.data[index] = index, tail

		delete(this.m, val)
		this.data = this.data[:len(this.data)-1]

		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
