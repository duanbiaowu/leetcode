package leetcode

import "testing"

func Test_candy(t *testing.T) {
	type args struct {
		ratings []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			"test-1",
			args{[]int{}},
			0,
		},
		{
			"test-2",
			args{[]int{1, 0, 2}},
			5,
		},
		{
			"test-3",
			args{[]int{1, 2, 2}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := candy(tt.args.ratings); gotAns != tt.wantAns {
				t.Errorf("candy() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
