package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRestaurant(t *testing.T) {
	type args struct {
		list1 []string
		list2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-1",
			args{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			},
			[]string{"Shogun"},
		},
		{
			"test-2",
			args{
				[]string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
				[]string{"KFC", "Shogun", "Burger King"},
			},
			[]string{"Shogun"},
		},
		{
			"test-3",
			args{
				[]string{"Tapioca Express", "Burger King", "KFC"},
				[]string{"KFC", "Shogun", "Burger King"},
			},
			[]string{"KFC"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRestaurant(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRestaurant() = %v, want %v", got, tt.want)
			}
		})
	}
}
