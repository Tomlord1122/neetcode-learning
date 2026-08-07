func findKthLargest(nums []int, k int) int {
	h := &maxHeap{}
	heap.Init(h)
	for _, num := range nums{
		heap.Push(h, num)
	}
	for i := 0; i < k-1; i++{
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

type maxHeap []int

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