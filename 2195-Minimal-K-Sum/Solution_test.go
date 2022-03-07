package leetcode

import "testing"

func Test_minimalKSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		//{
		//	"test-1",
		//	args{nil, 0},
		//	0,
		//},
		//{
		//	"test-2",
		//	args{[]int{1}, 0},
		//	0,
		//},
		//{
		//	"test-3",
		//	args{[]int{1,4,25,10,25}, 2},
		//	5,
		//},
		//{
		//	"test-4",
		//	args{[]int{6, 7}, 6},
		//	23,
		//},
		//{
		//	"test-5",
		//	args{[]int{6, 7}, 5},
		//	15,
		//},
		{
			"test-6",
			args{[]int{2, 6}, 3},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalKSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minimalKSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
