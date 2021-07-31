package leetcode

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"test-1",
			args{[]byte{}},
			[]byte{},
		},
		{
			"test-2",
			args{[]byte{'a'}},
			[]byte{'a'},
		},
		{
			"test-3",
			args{[]byte{'a', 'a'}},
			[]byte{'a', 'a'},
		},
		{
			"test-4",
			args{[]byte{'a', 'b'}},
			[]byte{'b', 'a'},
		},
		{
			"test-5",
			args{[]byte{'a', 'b', 'c'}},
			[]byte{'c', 'b', 'a'},
		},
		{
			"test-6",
			args{[]byte{'h', 'e', 'l', 'l', 'o'}},
			[]byte{'o', 'l', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
