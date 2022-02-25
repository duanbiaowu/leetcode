package leetcode

func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if k >= n {
		return arr
	}
	return quickSort(arr, k, 0, n-1)
}

func quickSort(arr []int, k, left, right int) []int {
	pivot := arr[right]
	i, j := left, right
	for i < j {
		for i < j && arr[i] <= pivot {
			i++
		}
		for i < j && arr[j] >= pivot {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[right] = arr[right], arr[i]

	if i > k {
		return quickSort(arr, k, left, i-1)
	}
	if i < k {
		return quickSort(arr, k, i+1, right)
	}
	return arr[:k]
}
