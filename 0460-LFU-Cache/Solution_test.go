package leetcode

import (
	"testing"
)

func TestLFUCache_Get(t *testing.T) {
	cache := Constructor(2)

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	// cache=[1,_], cnt(1)=1
	cache.Put(1, 1)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	// cache=[2,1], cnt(2)=1, cnt(1)=1
	cache.Put(2, 2)
	if got := cache.Get(2); got != 2 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	// 去除键 1 ，因为 cnt(1)=1 ，使用计数最小
	cache.Put(3, 3)
	if got := cache.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	// cache=[3,2], cnt(3)=2, cnt(2)=2

	// 去除键 2 ，2 和 3 的 cnt 相同，但 2 最久未使用
	// cache=[4,3], cnt(4)=1, cnt(3)=2
	cache.Put(4, 4)
	if got := cache.Get(4); got != 4 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(2); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	// cache=[4,3], cnt(4)=2, cnt(3)=3
}
