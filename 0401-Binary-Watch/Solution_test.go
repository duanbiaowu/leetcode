package leetcode

import (
	"reflect"
	"testing"
)

func Test_findNthDigit(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{1},
			[]string{"0:01", "0:02", "0:04", "0:08", "0:16", "0:32", "1:00", "2:00", "4:00", "8:00"},
		},
		{
			"test-2",
			args{9},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readBinaryWatch(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNthDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}
