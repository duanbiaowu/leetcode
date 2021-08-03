package leetcode

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	cycleListNode := &ListNode{Val: 1}
	cycleListNode.Next = &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: cycleListNode,
			},
		},
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{nil},
			false,
		},
		{
			"test-2",
			args{&ListNode{}},
			false,
		},
		{
			"test-3",
			args{&ListNode{Val: 1}},
			false,
		},
		{
			"test-4",
			args{&ListNode{
				Val:  1,
				Next: &ListNode{Val: 2},
			}},
			false,
		},
		{
			"test-5",
			args{cycleListNode},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
