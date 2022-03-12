package leetcode

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&Node{Val: 1}},
			[]int{1},
		},
		{
			"test-3",
			args{&Node{
				Val: 1,
				Children: []*Node{
					{
						Val: 3,
						Children: []*Node{
							{Val: 5},
							{Val: 6},
						},
					},
					{Val: 2},
					{Val: 4},
				},
			}},
			[]int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
