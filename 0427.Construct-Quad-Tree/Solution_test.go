package leetcode

import (
	"reflect"
	"testing"
)

func Test_construct(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *QuadTreeNode
	}{
		{
			"test-1",
			args{[][]int{
				{0, 1},
				{1, 0},
			}},
			&QuadTreeNode{
				IsLeaf: false,
				Val:    true,
				TopLeft: &QuadTreeNode{
					IsLeaf: true,
					Val:    false,
				},
				TopRight: &QuadTreeNode{
					IsLeaf: true,
					Val:    true,
				},
				BottomLeft: &QuadTreeNode{
					IsLeaf: true,
					Val:    true,
				},
				BottomRight: &QuadTreeNode{
					IsLeaf: true,
					Val:    false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := construct(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("construct() = %v, want %v", got, tt.want)
			}
		})
	}
}
