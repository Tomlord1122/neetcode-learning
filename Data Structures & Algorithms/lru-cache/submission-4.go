type Node struct{
	key, value int
	prev, next *Node
}
type LRUCache struct {
    capacity int
	cache map[int]*Node
	left, right *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		left: &Node{},
		right: &Node{},
	}
	lru.left.next, lru.right.prev = lru.right, lru.left
	return lru
}

func (this *LRUCache) remove(node *Node){
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *Node){
	// insert to the mfu frequency
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev
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
    if node, exist := this.cache[key]; exist{
		this.remove(node)
		delete(this.cache, key)
	}
	node := &Node{key:key, value:value}
	this.cache[key] = node
	this.insert(node)
	if len(this.cache) > this.capacity{
		// remove left.next
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}
}


// HashMap -> doubly linked list -> insert and remove