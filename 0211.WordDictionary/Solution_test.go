package leetcode

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()

	if got := obj.Search("hello"); got != false {
		t.Errorf("Get() = %v, want %v", got, false)
	}

	obj.AddWord("hello")

	if got := obj.Search("hello"); got != true {
		t.Errorf("Get() = %v, want %v", got, true)
	}
}
