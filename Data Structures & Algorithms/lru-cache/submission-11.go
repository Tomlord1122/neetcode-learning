type node struct{
	key int
	value int
	prev, next *node
}
type LRUCache struct {
	capacity int
	cache map[int]*node
	left, right *node
}

func Constructor(capacity int) LRUCache {
    lru := LRUCache{
		capacity: capacity,
		cache: make(map[int]*node),
		left: &node{},
		right: &node{},
	}
	lru.left.next, lru.right.prev = lru.right, lru.left
	return lru
}

func (this *LRUCache) insert(node *node){
	prev, next := this.right.prev, this.right
	node.prev = prev
	node.next = next
	prev.next = node
	next.prev = node
}

func (this *LRUCache) remove(node *node){
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.cache[key]; exist{
		// update the node and return value
		this.remove(node)
		this.insert(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exist := this.cache[key]; exist{
		this.remove(node)
	}
	node := &node{key:key, value:value}
	this.cache[key] = node
	this.insert(node)
	if len(this.cache) > this.capacity{
		// remove lru node
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}
