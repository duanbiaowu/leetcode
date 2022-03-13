package leetcode

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	index := make(map[string]int, len(list1))
	for i := range list1 {
		index[list1[i]] = i
	}

	indexSum := math.MaxInt32
	for i := range list2 {
		if j, ok := index[list2[i]]; ok {
			if i+j < indexSum {
				indexSum = i + j
				res = []string{list2[i]}
			} else if i+j == indexSum {
				res = append(res, list2[i])
			}
		}
	}

	return res
}
