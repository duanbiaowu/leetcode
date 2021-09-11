package leetcode

import (
	"reflect"
	"testing"
)

func Test_cloneGraph(t *testing.T) {
	type args struct {
		node *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"test-1",
			args{nil},
			nil,
		},
		{
			"test-2",
			args{&Node{1, []*Node{}}},
			&Node{1, []*Node{}},
		},
		{
			"test-3",
			args{&Node{1, []*Node{
				{2, []*Node{
					{3, []*Node{}},
					{4, []*Node{}},
				}},
			}}},
			&Node{1, []*Node{
				{2, []*Node{
					{3, []*Node{}},
					{4, []*Node{}},
				}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cloneGraph(tt.args.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}
