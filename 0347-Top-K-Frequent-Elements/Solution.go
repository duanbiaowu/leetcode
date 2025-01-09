package leetcode

// 自定义数字+频率对象
type node struct {
	num int
	cnt int
}

// 大顶堆实现方法

// 存在不足: 构建大顶堆的空间复杂度为 O(N), 如果数组很大，而 k 很小，会造成不必要的内存空间浪费
// 这是由大顶堆的性质决定的，因为堆顶始终是当前堆中频率最高的元素。
// 所以当需要获取前 K 个高频元素时，大顶堆会先存储所有频率元素对，
// 然后通过弹出 m−K 个元素（即频率较低的部分）来获取剩余的 K 个频率最高的元素。

// 改进方法: 通过让堆的大小始终保持在 K，可以在构建过程中节省内存空间并提高效率。
// 具体实现上，我们需要使用小顶堆而不是大顶堆，这样在遍历频率表时可以直接剔除低频元素，
// 确保堆中保留的是频率最高的 K 个元素
func topKFrequentMaxHeap(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}

	// 构建计数器 Map
	numCnt := make(map[int]int)
	for _, v := range nums {
		numCnt[v]++
	}

	heap := []*node{}
	for num, cnt := range numCnt {
		heap = append(heap, &node{
			num: num,
			cnt: cnt,
		})
	}

	// 构建大顶堆
	size := len(heap)
	buildMaxHeap(heap, size)

	res := []int{}

	for i := 0; i < k; i++ {
		// 堆顶就是当前频率最大的元素
		res = append(res, heap[0].num)

		heap[0], heap[size-1] = heap[size-1], heap[0]
		size--
		maxHeapify(heap, 0, size)
	}

	return res
}

func buildMaxHeap(arr []*node, size int) {
	// 从最后一个非叶子节点开始调整
	for i := size>>1 - 1; i >= 0; i-- {
		maxHeapify(arr, i, size)
	}
}

func maxHeapify(arr []*node, i, size int) {
	// 初始化最大值为根节点
	largest := i
	left, right := i*2+1, i*2+2

	if left < size && arr[left].cnt > arr[largest].cnt {
		largest = left
	}
	if right < size && arr[right].cnt > arr[largest].cnt {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, size)
	}
}

// 小顶堆实现方法

// 相较于大顶堆，小顶堆的优势在于:
// 小顶堆的堆顶始终是当前堆中频率最低的元素。
// 这样，当我们遍历频率表时，新的元素如果比堆顶的元素频率高，可以直接替换堆顶并调整堆。
// 在遍历过程中，当堆大小超过 K 时，我们只需要移除堆顶即可，始终保留频率最高的 K 个元素即可。
// 小顶堆是以淘汰最低频率元素为核心逻辑，其调整操作（插入、删除）正好完全符合题目需求。
func topKFrequentMinHeap(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}

	// 构建计数器 Map
	numCnt := make(map[int]int)
	for _, v := range nums {
		numCnt[v]++
	}

	heap := []*node{}
	for num, cnt := range numCnt {
		addToHeap(&heap, k, &node{num, cnt})
	}

	res := []int{}
	for _, v := range heap {
		res = append(res, v.num)
	}

	return res
}

// 将元素添加到小顶堆
func addToHeap(arr *[]*node, k int, v *node) {
	*arr = append(*arr, v)

	i := len(*arr) - 1

	// 小顶堆的上浮操作
	for i > 0 {
		parent := (i - 1) >> 1
		if (*arr)[i].cnt >= (*arr)[parent].cnt {
			break
		}

		(*arr)[i], (*arr)[parent] = (*arr)[parent], (*arr)[i]
		i = parent
	}

	// 如果堆的大小超过 k，移除堆顶元素
	if len(*arr) > k {
		size := len(*arr) - 1

		// 将对顶元素和最后一个元素交换
		(*arr)[0], (*arr)[size] = (*arr)[size], (*arr)[0]
		*arr = (*arr)[:size]

		minHeapify(*arr, 0, len(*arr))
	}
}

// 小顶堆的 minHeapify 操作
func minHeapify(arr []*node, i, size int) {
	// 初始化最小值为根节点
	smallest := i
	left, right := 2*i+1, 2*i+2

	// 如果左子节点更小
	if left < size && arr[left].cnt < arr[smallest].cnt {
		smallest = left
	}

	// 如果右子节点更小
	if right < size && arr[right].cnt < arr[smallest].cnt {
		smallest = right
	}

	// 如果最小值不是根节点
	if smallest != i {
		arr[i], arr[smallest] = arr[smallest], arr[i]
		minHeapify(arr, smallest, size)
	}
}

// 桶排序实现方法

func topKFrequent2(nums []int, k int) []int {
	if len(nums) < k {
		return nums
	}

	// 统计每个数字出现的次数
	// key: 数字, value: 出现的次数
	// 例如: nums = [1, 1, 1, 2, 2, 3, 4, 5],
	// numMap = {1: 3, 2: 2, 3: 1, 4: 1, 5: 1}
	numMap := make(map[int]int)
	for i := range nums {
		numMap[nums[i]]++
	}

	// 按照数字出现的次数进行桶排序
	// key: 出现的次数, value: 出现次数为 key 的数字列表
	// 例如: cntMap = {3: [1], 2: [2], 1: [3, 4, 5]}
	cntMap := make(map[int][]int)
	for num, cnt := range numMap {
		cntMap[cnt] = append(cntMap[cnt], num)
	}

	res := []int{}

	// 题目要求: 1 <= nums.length <= 105
	// 从后往前遍历, 保证出现次数多的数字在前面
	for i := len(nums); i > 0; i-- {
		if _, ok := cntMap[i]; !ok {
			continue
		}

		for _, num := range cntMap[i] {
			res = append(res, num)
			// 题目要求返回前 k 个出现次数最多的数字
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
