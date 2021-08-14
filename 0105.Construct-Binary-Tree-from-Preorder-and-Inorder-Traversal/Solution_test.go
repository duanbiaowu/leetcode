package leetcode

import (
	"github.com/duanbiaowu/leetcode/structures"
	"reflect"
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{nil, nil},
			nil,
		},
		{
			"test-2",
			args{[]int{}, []int{}},
			nil,
		},
		{
			"test-3",
			args{[]int{1}, []int{1}},
			&TreeNode{Val: 1},
		},
		{
			"test-4",
			args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			structures.GenerateTreeNodesBySlice([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.preorder, tt.args.inorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
