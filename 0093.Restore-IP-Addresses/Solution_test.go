package leetcode

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{""},
			nil,
		},
		{
			"test-2",
			args{"25525511135"},
			[]string{
				"255.255.111.35",
				"255.255.11.135",
			},
		},
		{
			"test-3",
			args{"0000"},
			[]string{"0.0.0.0"},
		},
		{
			"test-4",
			args{"1111"},
			[]string{"1.1.1.1"},
		},
		{
			"test-5",
			args{"010010"},
			[]string{
				"0.100.1.0",
				"0.10.0.10",
			},
		},
		{
			"test-6",
			args{"101023"},
			[]string{
				"101.0.2.3",
				"10.10.2.3",
				"10.1.0.23",
				"1.0.102.3",
				"1.0.10.23",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
