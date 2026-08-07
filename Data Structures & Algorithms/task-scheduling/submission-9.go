func leastInterval(tasks []byte, n int) int {
	freqMap := make(map[byte]int)
	for _, c := range tasks{
		freqMap[c]++
	}
	h := &MaxHeap{}
	heap.Init(h)
	for _, f := range freqMap{
		heap.Push(h, f)
	}

	queue := [][]int{} // (remain freq, delay time)
	time := 0
	for h.Len() > 0 || len(queue) > 0{
		time++
		if len(queue) > 0 && time >= queue[0][1]{
			cur := queue[0][0]
			queue = queue[1:]
			heap.Push(h, cur)
		}
		if h.Len() > 0{
			pop := heap.Pop(h).(int)
			if pop - 1 > 0{
			queue = append(queue, []int{pop-1, time+n+1})
			}
		}
	}
	return time
}

// calculate the frequency first
// then use a maxHeap -> to pop this
// maintain a queue to track the delayed character
// every single turn we should use a variable to track the time which start at zero

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
