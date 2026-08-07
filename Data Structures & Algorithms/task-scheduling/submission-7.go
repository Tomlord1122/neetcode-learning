func leastInterval(tasks []byte, n int) int {
	numFreq := make(map[byte]int) // num, freq
	for _, task := range tasks{
		numFreq[task]++
	}

	h := &maxHeap{}
	heap.Init(h)
	for _, freq := range numFreq{
		heap.Push(h, freq)
	}
	queue := [][]int{} // [0] -> time [1] -> value

	time := 0
	for h.Len() > 0 || len(queue) > 0{
		time++
		// check if we can pop something from queue
		if len(queue) > 0 && time >= queue[0][0]{
			cur := queue[0]
			queue = queue[1:]
			// push cur to heap
			heap.Push(h, cur[1])
		}
		if h.Len() > 0{
			cur := heap.Pop(h).(int)
			if cur - 1 > 0{
				queue = append(queue, []int{time+n+1, cur-1})
			}
		}
	}
	return time
}



type maxHeap []int

// Push, Pop, Less, Swap, Len
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

func (h maxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Len() int{
	return len(h)
}