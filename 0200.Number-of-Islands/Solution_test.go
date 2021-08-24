package leetcode

import (
	"testing"
)

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{[][]byte{}},
			0,
		},
		{
			"test-2",
			args{[][]byte{
				{'0'},
			}},
			0,
		},
		{
			"test-3",
			args{[][]byte{
				{'1'},
			}},
			1,
		},
		{
			"test-4",
			args{[][]byte{
				{'0', '0'},
				{'0', '0'},
			}},
			0,
		},
		{
			"test-5",
			args{[][]byte{
				{'1', '1'},
				{'1', '1'},
			}},
			1,
		},
		{
			"test-6",
			args{[][]byte{
				{'1', '0'},
				{'0', '1'},
			}},
			2,
		},
		{
			"test-7",
			args{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
