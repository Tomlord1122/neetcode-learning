func networkDelayTime(times [][]int, n int, k int) int {
	// create graph via adjacency list
	edgeMap := make(map[int][]Edge)
	for _, time := range times{
		u, v, w := time[0], time[1], time[2]
		edgeMap[u] = append(edgeMap[u], Edge{node: v, weight: w})
	}
	// create the minHeap and push node k into it
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Edge{node:k, weight:0})

	visited := make(map[int]bool)
	t := 0
	// Run the dijkstra algorithm
	for h.Len() > 0{
		edge := heap.Pop(h).(Edge)
		node, time := edge.node, edge.weight
		if visited[node]{
			continue
		}
		visited[node] = true
		t = time
		for _, next := range edgeMap[node]{
			if visited[next.node]{
				continue
			}
			heap.Push(h, Edge{node:next.node, weight: time+next.weight})
		}
	}
	// check the condition that we've already visit all the node
	if len(visited) == n{
		return t
	}
	return -1
}

type Edge struct{
	node int
	weight int
}

type MinHeap []Edge

// Push, Pop, Less, Len, Swap
func (h *MinHeap) Push(x any){
	*h = append(*h, x.(Edge))
}

func (h *MinHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func (h MinHeap) Less(i, j int) bool{
	return h[i].weight < h[j].weight
}

func (h MinHeap) Len() int{
	return len(h)
}

func (h MinHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}
