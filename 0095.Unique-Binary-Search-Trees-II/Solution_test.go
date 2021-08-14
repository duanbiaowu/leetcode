package leetcode

import (
	"reflect"
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			"test-1",
			args{0},
			[]*TreeNode{nil},
		},
		{
			"test-1",
			args{1},
			[]*TreeNode{
				{Val: 1},
			},
		},
		{
			"test-2",
			args{2},
			[]*TreeNode{
				{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateTrees(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
