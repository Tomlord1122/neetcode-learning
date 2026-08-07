func lastStoneWeight(stones []int) int {
    h := &maxHeap{}
    heap.Init(h)

    for _, stone := range stones{
        heap.Push(h, stone)
    }

    for h.Len() > 1{
        a, b := heap.Pop(h).(int), heap.Pop(h).(int)
        if a == b{
            continue
        } else {
            heap.Push(h, a - b)
        }
    }
    if h.Len() == 0{
        return 0
    }
    return (*h)[0]
}

// heap/container
// Push, Pop, Len, Swap, Less
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

func (h maxHeap) Swap(i, j int){
    h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool{
    return h[i] > h[j]
}