package leetcode

import "testing"

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"test-1",
			args{0},
			0,
		},
		{
			"test-2",
			args{43261596},
			964176192,
		},
		{
			"test-3",
			args{4294967293},
			3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBits(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
