func kClosest(points [][]int, k int) [][]int {
	h := &minHeap{}
	heap.Init(h)
	for _, point := range points{
		distance := math.Sqrt(float64(point[0]*point[0]) + float64(point[1]*point[1]))
		heap.Push(h, pair{coordinate: point, distance:distance})
	}
	res := [][]int{}
	for i := 0; i < k; i++{
		ans := heap.Pop(h).(pair)
		res = append(res, ans.coordinate)
	}
	return res
}

type pair struct{
	coordinate []int
	distance float64
}

type minHeap []pair

func (h *minHeap) Push(x any){
	*h = append(*h, x.(pair))
}

func (h *minHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h minHeap) Less(i, j int) bool{
	return h[i].distance < h[j].distance
}

func (h minHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Len() int{
	return len(h)
}