package leetcode

import "testing"

func Test_numberToWords(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test-1",
			args{1},
			"One",
		},
		{
			"test-2",
			args{10},
			"Ten",
		},
		{
			"test-3",
			args{123},
			"One Hundred Twenty Three",
		},
		{
			"test-4",
			args{12345},
			"Twelve Thousand Three Hundred Forty Five",
		},
		{
			"test-5",
			args{1234567},
			"One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			"test-6",
			args{1234567891},
			"One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberToWords(tt.args.num); got != tt.want {
				t.Errorf("numberToWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
