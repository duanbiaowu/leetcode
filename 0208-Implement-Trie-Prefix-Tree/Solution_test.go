package leetcode

import "testing"

func TestTrie(t *testing.T) {
	obj := Constructor()

	if got := obj.Search("hello"); got != false {
		t.Errorf("Get() = %v, want %v", got, false)
	}
	if got := obj.StartsWith("hello"); got != false {
		t.Errorf("Get() = %v, want %v", got, false)
	}

	obj.Insert("hello")

	if got := obj.Search("hello"); got != true {
		t.Errorf("Get() = %v, want %v", got, true)
	}
	if got := obj.StartsWith("hello"); got != true {
		t.Errorf("Get() = %v, want %v", got, true)
	}
	if got := obj.Search("he"); got != false {
		t.Errorf("Get() = %v, want %v", got, false)
	}
	if got := obj.StartsWith("he"); got != true {
		t.Errorf("Get() = %v, want %v", got, true)
	}

	obj.Insert("he")
	if got := obj.Search("he"); got != true {
		t.Errorf("Get() = %v, want %v", got, true)
	}
}
