package leetcode

import (
	"testing"
)

func Test_respace(t *testing.T) {
	type args struct {
		dictionary []string
		sentence   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test-1",
			args{nil, ""},
			0,
		},
		{
			"test-2",
			args{[]string{"abc"}, "d"},
			1,
		},
		{
			"test-3",
			args{[]string{"looked", "just", "like", "her", "brother"}, "jesslookedjustliketimherbrother"},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := respace(tt.args.dictionary, tt.args.sentence); got != tt.want {
				t.Errorf("respace() = %v, want %v", got, tt.want)
			}
		})
	}
}
