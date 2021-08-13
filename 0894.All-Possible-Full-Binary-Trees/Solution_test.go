package leetcode

import (
	"reflect"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
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
			nil,
		},
		{
			"test-2",
			args{1},
			[]*TreeNode{
				{Val: 0},
			},
		},
		{
			"test-3",
			args{3},
			[]*TreeNode{
				{
					Val:   0,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
