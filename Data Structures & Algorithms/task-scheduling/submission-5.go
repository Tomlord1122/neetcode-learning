func leastInterval(tasks []byte, n int) int {
	taskCount := make(map[byte]int)
	for _, task := range tasks{
		taskCount[task]++
	}
	h := &maxHeap{}
	heap.Init(h)
	for _, v := range taskCount{
		heap.Push(h, v)
	}
	queue := []delay{}
	time := 0
	for h.Len() != 0 || len(queue) != 0{
		time++
		if len(queue) != 0 && queue[0].time == time{
			heap.Push(h, queue[0].count)
			queue = queue[1:]
		}
		if h.Len() == 0{
			continue
		}
		count := heap.Pop(h).(int)
		if count - 1 == 0{
			continue
		}
		queue = append(queue, delay{count:count-1, time:time+n+1})
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

func (h maxHeap) Less(i, j int) bool{
	return h[i] > h[j]
}

func (h maxHeap) Len() int{
	return len(h)
}

func (h maxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}
