package leetcode

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	listNodesWithoutCarry1 := &ListNode{Val: 1}
	listNodesWithoutCarry1.Next = &ListNode{Val: 3}
	listNodesWithoutCarry1.Next.Next = &ListNode{Val: 5}

	listNodesWithoutCarry2 := &ListNode{Val: 0}
	listNodesWithoutCarry2.Next = &ListNode{Val: 2}
	listNodesWithoutCarry2.Next.Next = &ListNode{Val: 4}

	listNodesWithoutCarry3 := &ListNode{Val: 1}
	listNodesWithoutCarry3.Next = &ListNode{Val: 5}
	listNodesWithoutCarry3.Next.Next = &ListNode{Val: 9}

	listNodesWithCarry1 := &ListNode{Val: 2}
	listNodesWithCarry1.Next = &ListNode{Val: 4}
	listNodesWithCarry1.Next.Next = &ListNode{Val: 3}

	listNodesWithCarry2 := &ListNode{Val: 5}
	listNodesWithCarry2.Next = &ListNode{Val: 6}
	listNodesWithCarry2.Next.Next = &ListNode{Val: 4}

	listNodesWithCarry3 := &ListNode{Val: 7}
	listNodesWithCarry3.Next = &ListNode{Val: 0}
	listNodesWithCarry3.Next.Next = &ListNode{Val: 8}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test two empty linkNodes",
			args{nil, nil},
			nil,
		},
		{
			"test single empty linkNode-1",
			args{&ListNode{Val: 1}, nil},
			&ListNode{Val: 1},
		},
		{
			"test single empty linkNode-2",
			args{nil, &ListNode{Val: 2}},
			&ListNode{Val: 2},
		},
		{
			"test two linkNodes with only one element",
			args{&ListNode{Val: 1}, &ListNode{Val: 2}},
			&ListNode{Val: 3},
		},
		{
			"test two linkNodes with multiple element, but there is no carry",
			args{listNodesWithoutCarry1, listNodesWithoutCarry2},
			listNodesWithoutCarry3,
		},
		{
			"test two linkNodes with multiple element, and there is carry",
			args{listNodesWithCarry1, listNodesWithCarry2},
			listNodesWithCarry3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
