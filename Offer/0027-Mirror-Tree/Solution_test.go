package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_mirrorTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&TreeNode{}},
			&TreeNode{},
		},
		{
			"test-3",
			args{structures.GenerateTreeNodesBySlice([]int{4, 2, 7, 1, 3, 6, 9})},
			structures.GenerateTreeNodesBySlice([]int{4, 7, 2, 9, 6, 3, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mirrorTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
