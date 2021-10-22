package leetocde

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&ListNode{}},
			&ListNode{},
		},
		{
			"test-3",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4})},
			structures.GenerateListNodesByArray([]int{2, 1, 4, 3}),
		},
		{
			"test-4",
			args{structures.GenerateListNodesByArray([]int{1, 2, 3, 4, 5})},
			structures.GenerateListNodesByArray([]int{2, 1, 4, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
