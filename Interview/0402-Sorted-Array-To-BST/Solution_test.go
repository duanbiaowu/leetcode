package leetcode

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			"test-1",
			args{[]int{}},
			nil,
		},
		{
			"test-2",
			args{[]int{1}},
			&TreeNode{Val: 1},
		},
		{
			"test-3",
			args{[]int{1, 2, 3}},
			&TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
