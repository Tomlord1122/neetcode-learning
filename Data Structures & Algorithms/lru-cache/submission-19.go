type Node struct{
	key int
	value int
	prev, next *Node
}
type LRUCache struct {
    capacity int
	cache map[int]*Node
	left, right *Node
}

func Constructor(capacity int) LRUCache {
    left, right := &Node{}, &Node{}
	left.next, right.prev = right, left
	return LRUCache{
		capacity: capacity,
		cache: make(map[int]*Node),
		left: left,
		right: right,
	}
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
	newNode := &Node{key:key, value: value}
	this.cache[key] = newNode
	this.Insert(newNode)
	if len(this.cache) > this.capacity{
		// remove lru node
		lru := this.left.next
		this.Remove(lru)
		delete(this.cache, lru.key)
	}
}

func (this *LRUCache) Insert(node *Node){
	prev, next := this.right.prev, this.right
	node.next = next
	node.prev = prev
	prev.next = node
	next.prev = node
}

func (this *LRUCache) Remove(node *Node){
	prev, next := node.prev, node.next
	prev.next, next.prev = next, prev
}