package leetcode

import "testing"

func TestLRUCache_Get(t *testing.T) {
	cache := Constructor(2)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	cache.Put(1, 1)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get() = %v, want %v", got, 1)
	}

	cache.Put(2, 2)
	if got := cache.Get(2); got != 2 {
		t.Errorf("Get() = %v, want %v", got, 2)
	}

	cache.Put(3, 3)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(3); got != 3 {
		t.Errorf("Get() = %v, want %v", got, 3)
	}

	cache.Put(4, 4)
	if got := cache.Get(2); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(4); got != 4 {
		t.Errorf("Get() = %v, want %v", got, 4)
	}

	cache.Put(5, 5)
	cache.Put(6, 6)
	cache.Put(7, 7)
	cache.Put(8, 8)

	if got := cache.Get(5); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}
	if got := cache.Get(6); got != -1 {
		t.Errorf("Get() = %v, want %v", got, -1)
	}

	if got := cache.Get(7); got != 7 {
		t.Errorf("Get() = %v, want %v", got, 7)
	}
	if got := cache.Get(8); got != 8 {
		t.Errorf("Get() = %v, want %v", got, 8)
	}
}
