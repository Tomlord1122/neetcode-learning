func findKthLargest(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range nums{
		heap.Push(h, num)
	}

	for i := 0; i < k-1; i++{
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}


type MaxHeap []int

// Push, Pop, Len, Less, Swap
func (h *MaxHeap) Push(x any){
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h MaxHeap) Len() int{
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}
