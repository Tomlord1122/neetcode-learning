func lastStoneWeight(stones []int) int {
	h := &maxHeap{}
	heap.Init(h)
	for _, stone := range stones{
		heap.Push(h, stone)
	}
	for h.Len() >= 2{
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x == y{
			continue
		}
		heap.Push(h, x - y)
	}

	if h.Len() == 1{
		return (*h)[0]
	}
	return 0
}

type maxHeap []int

// Push, Pop, Less, Len, Swap
func (h *maxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h maxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h maxHeap) Len() int{
	return len(h)
}

func (h maxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}