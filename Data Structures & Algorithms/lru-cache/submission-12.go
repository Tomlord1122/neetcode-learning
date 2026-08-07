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
    l, r := &node{}, &node{}
    l.next, r.prev = r, l
    return LRUCache{
		capacity: capacity,
		cache: make(map[int]*node),
		left: l,
		right: r,
	}
}

func (this *LRUCache) Insert(node *node){
	prev, next := this.right.prev, this.right
	node.prev = prev
	node.next = next
	prev.next = node
	next.prev = node
}

func (this *LRUCache) Remove(node *node){
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) Get(key int) int {
	if node, exist := this.cache[key]; exist{
		this.Remove(node)
		this.Insert(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, exist := this.cache[key]; exist{
		this.Remove(node)
	}
	newNode := &node{key:key, value:value}
	this.cache[key] = newNode
	this.Insert(newNode)

	if len(this.cache) > this.capacity{
		lru := this.left.next
		this.Remove(lru)
		delete(this.cache, lru.key)
	}
}