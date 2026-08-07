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

func (this *LRUCache) remove(node *node){
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *node){
	// Add it to right (MFU)
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.prev = prev
	node.next = next
}

func (this *LRUCache) Get(key int) int {
    if node, exist := this.cache[key]; exist{
		this.remove(node)
		this.insert(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
    // if key is exist
	if node, exist := this.cache[key]; exist{
		this.remove(node)
		delete(this.cache, key)
	}
	// add node and check capacity 
	node := &node{key:key, value:value}
	this.cache[key] = node
	this.insert(node)
	if len(this.cache) > this.capacity{
		// remove left
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}


// Doubly linked list + HashMap
// Insert node to linked list
// remove node to linked list