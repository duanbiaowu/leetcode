package leetcode

import "testing"

func Test_dayOfYear(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{"2022-01-01"},
			1,
		},
		{
			"test-2",
			args{"2022-01-09"},
			9,
		},
		{
			"test-3",
			args{"2022-02-10"},
			41,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dayOfYear(tt.args.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
