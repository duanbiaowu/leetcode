package leetcode

// reference: https://leetcode-cn.com/problems/lfu-cache/solution/chao-xiang-xi-tu-jie-dong-tu-yan-shi-460-lfuhuan-c/

// LRU(Least Recently Used)
// 最近最少使用算法，它是根据时间维度来选择将要淘汰的元素，
// 即删除掉最长时间没被访问的元素。

// LFU(Least Frequently Used)
// 最近最不常用算法，它是根据频率维度来选择将要淘汰的元素，
// 即删除访问频率最低的元素。
// 如果两个元素的访问频率相同，则淘汰最久没被访问的元素。
// 也就是说LFU淘汰的时候会选择两个维度，先比较频率，选择访问频率最小的元素；
// 如果频率相同，则按时间维度淘汰掉最久远的那个元素。

// 双向链表节点
type Node struct {
	key  int
	val  int
	freq int
	pre  *Node
	next *Node
}

func NewNode(key, val, freq int) *Node {
	return &Node{key: key, val: val, freq: freq}
}

// 双向链表
type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	list := &LinkedList{
		head: NewNode(-1, -1, -1),
		tail: NewNode(-1, -1, -1),
	}
	list.head.next = list.tail
	list.tail.next = list.head
	return list
}

func (list *LinkedList) addToHead(node *Node) {
	if node != nil {
		node.pre = list.head
		node.next = list.head.next
		list.head.next.pre = node
		list.head.next = node
	}
}

func (list *LinkedList) removeNode(node *Node) {
	if node != nil {
		node.pre.next = node.next
		node.next.pre = node.pre
		node.pre, node.next = nil, nil
	}
}

func (list *LinkedList) getLastNode() *Node {
	if list.isEmpty() {
		return NewNode(-1, -1, -1)
	}
	return list.tail.pre
}

func (list *LinkedList) isEmpty() bool {
	return list.head.next == list.tail
}

// minFreq 用来记录 LFU 缓存中频率最小的元素，
// 在缓存满的时候，可以快速定位到最小频繁的链表，
// 以达到 O(1) 时间复杂度删除一个元素
// 具体的逻辑：
// 	更新/查找的时候，将元素频率 +1，之后如果 minFreq 不在频率哈希表中了，
//		说明频率哈希表中已经没有元素了，那么 minFreq 需要+1，否则 minFreq 不变。
// 	插入的时候，因为新元素的频率都是 1，所以只需要将 minFreq 改为 1 即可。
// minFreq 不在意具体的频率顺序，只需要保证指向最小的频率就 OK

type LFUCache struct {
	// key -> value
	keyMap map[int]*Node
	// key -> LinkedList
	freqMap  map[int]*LinkedList
	capacity int
	minFreq  int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		make(map[int]*Node),
		make(map[int]*LinkedList),
		capacity,
		0,
	}
}

// 如果key不存在则返回 -1
// 如果key存在，则返回对应的value，同时:
// 	将元素的访问频率+1
// 	将元素从访问频率i的链表中移除，放到频率i+1的链表中
// 	如果频率i的链表为空，则从频率哈希表中移除这个链表

func (this *LFUCache) Get(key int) int {
	if node, ok := this.keyMap[key]; ok {
		this.increment(node, false)
		return node.val
	}
	return -1
}

// 如果 key 已经存在，修改对应的 value，并将访问频率 +1
// 	将元素从访问频率i的链表中移除，放到频率 i+1 的链表中
//	如果频率 i 的链表为空，则从频率哈希表中移除这个链表
// 如果 key 不存在
// 	缓存超过最大容量，则先删除访问频率最低的元素，再插入新元素
// 		新元素的访问频率为1，如果频率哈希表中不存在对应的链表需要创建
//	缓存没有超过最大容量，则插入新元素
//		新元素的访问频率为1，如果频率哈希表中不存在对应的链表需要创建

func (this *LFUCache) Put(key int, value int) {
	if node, ok := this.keyMap[key]; ok {
		node.val = value
		this.increment(node, false)
	} else {
		if this.capacity == 0 {
			return
		}
		if len(this.keyMap) == this.capacity {
			this.remoteMinFreqNode()
		}
		node = NewNode(key, value, 1)
		this.increment(node, true)
		this.keyMap[key] = node
	}
}

// 更新节点的访问频率
// isNewNode 是否是新节点，新插入的节点和非新插入节点更新逻辑不同
func (this *LFUCache) increment(node *Node, isNewNode bool) {
	if isNewNode {
		this.minFreq = 1
		this.addToLinkedList(node)
	} else {
		this.removeNode(node)
		node.freq++
		this.addToLinkedList(node)
		if _, ok := this.freqMap[this.minFreq]; !ok {
			this.minFreq++
		}
	}
}

// 根据节点的频率，插入到对应的LinkedList中
func (this *LFUCache) addToLinkedList(node *Node) {
	if _, ok := this.freqMap[node.freq]; !ok {
		this.freqMap[node.freq] = NewLinkedList()
	}
	this.freqMap[node.freq].addToHead(node)
}

// 删除指定的节点
// 如果节点删除后，对应的双链表为空，则从 freqMap 中删除这个链表
func (this *LFUCache) removeNode(node *Node) {
	linkedList := this.freqMap[node.freq]
	linkedList.removeNode(node)
	if linkedList.isEmpty() {
		delete(this.freqMap, node.freq)
	}
}

// 删除频率最低的元素
// 从 freqMap 和 keyMap 中都要删除这个节点
// 如果节点删除后，对应的双链表为空，则从 freqMap 中删除这个链表
func (this *LFUCache) remoteMinFreqNode() {
	linkedList := this.freqMap[this.minFreq]
	node := linkedList.getLastNode()

	this.freqMap[this.minFreq].removeNode(node)
	delete(this.keyMap, node.key)

	if linkedList.isEmpty() {
		delete(this.freqMap, this.minFreq)
	}
}
