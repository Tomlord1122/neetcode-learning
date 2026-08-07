func pickGifts(gifts []int, k int) int64 {
	h := &maxHeap{}
	heap.Init(h)
	
	for _, val := range gifts{
		heap.Push(h, val)
	}

	for i := 0; i < k; i++{
		tmp := heap.Pop(h).(int)
		cur := math.Floor(math.Sqrt(float64(tmp)))
		heap.Push(h, int(cur))
	}
	
	res := 0
	for i := 0; i < h.Len(); i++{
		res += (*h)[i]
	}
	return int64(res)
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



// In each iteration, we should get the max val -> maximum number of gifts
// Maybe we should use a maxHeap to get this val?
