func minCostConnectPoints(points [][]int) int {
	n := len(points)
	dsu := NewDSU(n)
	
	edges := [][]int{}
	res := 0
	for i := 0; i < n; i++{
		x1, y1 := points[i][0], points[i][1]
		for j := i+1; j < n; j++{
			x2, y2 := points[j][0], points[j][1]
			// Then we can start calcualte the dist between these two nodes
			dist := int(math.Abs(float64(x1 - x2)) + math.Abs(float64((y1 - y2))))
			edges = append(edges, []int{dist, i, j})
		}
	}

	sort.Slice(edges, func(i, j int)bool{
		return edges[i][0] < edges[j][0]
	})

	// we need n-1 edges to build the MST
	for _, edge := range edges{
		x, y := edge[1], edge[2]
		if dsu.Union(x, y){
			res += edge[0]
		}
	}
	return res
}


// use disjoint set to solve this problem
// the cost is using manhattan distance

// Here I want to use union-by-size
type DSU struct{
	Parent []int
	Size []int
}

func NewDSU(n int) *DSU{
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++{
		parent[i] = i
		size[i] = 1
	}
	return &DSU{
		Parent: parent,
		Size: size,
	}
}

// start implementing the find and union method
func (dsu *DSU) Find(x int) int{
	if dsu.Parent[x] != x{
		dsu.Parent[x] = dsu.Find(dsu.Parent[x])
	}
	return dsu.Parent[x]
}

func (dsu *DSU) Union(x, y int) bool{
	px, py := dsu.Find(x), dsu.Find(y)
	if px == py{
		return false
	}
	if dsu.Size[px] < dsu.Size[py]{
		px, py = py, px
	}
	// px's size is always greater than the py
	dsu.Parent[py] = px
	dsu.Size[px] += dsu.Size[py]
	return true
}