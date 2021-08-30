package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func TestBSTIterator_HasNext(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name   string
		fields args
		want   bool
	}{
		{
			"test-1",
			args{nil},
			false,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{2, 1, 3})},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := Constructor(tt.fields.root)
			if got := iterator.HasNext(); got != tt.want {
				t.Errorf("HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator_Next(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name   string
		fields args
		want   []int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{structures.GenerateTreeNodesBySlice([]int{
				7, 3, 15, structures.NULL, structures.NULL, 9, 20,
			})},
			[]int{3, 7, 9, 15, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			iterator := Constructor(tt.fields.root)

			for iterator.HasNext() {
				got = append(got, iterator.Next())
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
