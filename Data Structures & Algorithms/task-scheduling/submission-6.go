func leastInterval(tasks []byte, n int) int {
	taskCount := make(map[byte]int)
	for _, t := range tasks{
		taskCount[t]++
	}

	queue := []delay{}
	h := &maxHeap{}
	heap.Init(h)
	
	for _, v := range taskCount{
		heap.Push(h, v)
	}

	time := 0

	for h.Len() != 0 || len(queue) != 0{
		time++

		// Check if we can pop queue
		if len(queue) != 0 && queue[0].time == time{
			pop := queue[0]
			queue = queue[1:]
			heap.Push(h, pop.count)
		}

		if h.Len() != 0{
			pop := heap.Pop(h).(int)
			if pop - 1 == 0{
				continue
			} else {
				queue = append(queue, delay{count: pop-1, time: time+n+1})
			}
		}
	}
	return time
}

type delay struct{
	count int
	time int
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

func (h maxHeap) Len() int{
	return len(h)
}

func (h maxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}