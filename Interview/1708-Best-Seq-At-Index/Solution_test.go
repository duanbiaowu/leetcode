package leetcode

import "testing"

func Test_bestSeqAtIndex(t *testing.T) {
	type args struct {
		height []int
		weight []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, nil},
			0,
		},
		{
			"test-2",
			args{[]int{65, 70, 56, 75, 60, 68}, []int{100, 150, 90, 190, 95, 110}},
			6,
		},
		//{
		//	"test-3",
		//	args{
		//		[]int{2868,5485,1356,1306,6017,8941,7535,4941,6331,6181},
		//		[]int{5042,3995,7985,1651,5991,7036,9391,428,7561,8594},
		//	},
		//	5,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestSeqAtIndex(tt.args.height, tt.args.weight); got != tt.want {
				t.Errorf("bestSeqAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
