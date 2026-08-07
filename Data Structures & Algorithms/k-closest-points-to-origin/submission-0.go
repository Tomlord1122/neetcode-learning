func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)
	for _, point := range points{
		distance := math.Sqrt(float64(point[0] * point[0]) + float64(point[1] * point[1]))
		heap.Push(h, Pair{distance:distance, coordinate: point})
	}
	res := [][]int{}
	for i := 0; i < k; i++{
		v := heap.Pop(h).(Pair)
		res = append(res, v.coordinate)
	}
	return res
}


// return k closet points to the origin (0, 0)
type Pair struct{
	distance float64
	coordinate []int
}

type MinHeap []Pair

// Push, Pop, Len, Less, Swap
func (h *MinHeap) Push(x any){
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h MinHeap) Len() int{
	return len(h)
}

func (h MinHeap) Less(i, j int) bool{
	return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}