package solution

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
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
			args{&Node{}},
			&Node{},
		},
		{
			"test-3",
			args{&Node{
				Val: 1,
				Left: &Node{
					Val: 2,
				},
				Right: &Node{
					Val: 3,
				},
				Next: nil,
			}},
			&Node{
				Val: 1,
				Left: &Node{
					Val: 2,
					Next: &Node{
						Val: 3,
					},
				},
				Right: &Node{
					Val: 3,
				},
				Next: nil,
			},
		},
		{
			"test-4",
			args{&Node{
				Val: 1,

				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
					},
					Right: &Node{
						Val: 5,
					},
				},

				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},

				Next: nil,
			}},
			&Node{
				Val: 1,

				Left: &Node{
					Val: 2,
					Left: &Node{
						Val: 4,
						Next: &Node{
							Val: 5,
							Next: &Node{
								Val: 7,
							},
						},
					},
					Right: &Node{
						Val: 5,
						Next: &Node{
							Val: 7,
						},
					},
					Next: &Node{
						Val: 3,
						Right: &Node{
							Val: 7,
						},
					},
				},

				Right: &Node{
					Val: 3,
					Right: &Node{
						Val: 7,
					},
				},

				Next: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
