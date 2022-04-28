package leetcode

type LRUCache struct {
	cache     map[int]*DListNode
	head      *DListNode
	tail      *DListNode
	size, cap int
}

// 必须存key,在淘汰map中的元素时,是根据node中的key来淘汰的
type DListNode struct {
	key       int
	value     int
	pre, next *DListNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: make(map[int]*DListNode),
		head:  &DListNode{},
		tail:  &DListNode{},
		size:  0,
		cap:   capacity,
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.movToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.movToHead(node)
		return
	} else {
		n := &DListNode{
			key:   key,
			value: value,
		}
		this.size++
		this.addToHead(n)
		this.cache[key] = n
		if this.size > this.cap {
			tmp := this.removeTail()
			delete(this.cache, tmp.key)
			this.size--
		}
	}
}

func (this *LRUCache) movToHead(node *DListNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *DListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *DListNode) {
	node.next = this.head.next
	this.head.next.pre = node
	node.pre = this.head
	this.head.next = node
}

func (this *LRUCache) removeTail() *DListNode {
	node := this.tail.pre
	this.removeNode(node)
	return node
}
