func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, stone := range stones{
		heap.Push(h, stone)
	}
	for h.Len() > 1{
		x, y := heap.Pop(h).(int), heap.Pop(h).(int)
		if x == y{
			continue
		} else {
			heap.Push(h, x-y)
		}
	}
	if h.Len() == 1{
		return (*h)[0]
	}
	return 0
}

type MaxHeap []int

// Push, Pop, Less, Len, Swap
func (h *MaxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h MaxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h MaxHeap) Len() int{
	return len(h)
}

func (h MaxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

