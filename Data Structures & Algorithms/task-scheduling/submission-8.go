func leastInterval(tasks []byte, n int) int {
	freqMap := make(map[byte]int)
	for _, task := range tasks{
		freqMap[task]++
	}
	h := &maxHeap{}
	heap.Init(h)
	for _, freq := range freqMap{
		heap.Push(h, freq)
	}
	queue := [][]int{} // val, time
	time := 0
	for h.Len() > 0 || len(queue) > 0{
		time++
		if len(queue) > 0 && queue[0][1] <= time{
			// pop the queue and append it to heap
			pop := queue[0][0]
			queue = queue[1:]
			heap.Push(h, pop)
		}
		if h.Len() > 0 {
			cur := heap.Pop(h).(int)
			if cur - 1 > 0{
				queue = append(queue, []int{cur-1, time+n+1})
			}
		}
	}
	return time
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