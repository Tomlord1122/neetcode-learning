func findRedundantConnection(edges [][]int) []int {
    // use disjoint set data structure
	n := len(edges)
	parent := make([]int, n+1)
	height := make([]int, n+1)
	for i := 1; i <= n; i++{
		parent[i] = i
		height[i] = 1
	}
	// implement find and union function
	var find func(i int) int
	find = func(i int) int{
		if parent[i] != i{
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	union := func(x, y int) bool{
		rootX, rootY := find(x), find(y)
		if rootX == rootY{
			return false
		} else {
			if height[rootX] > height[rootY]{
				parent[rootY] = rootX
				height[rootX]++
			} else if height[rootX] < height[rootY]{
				parent[rootX] = rootY
				height[rootY]++
			} else {
				parent[rootY] = rootX
			}
		}
		return true		
	}

	for _, edge := range edges{
		if !union(edge[0], edge[1]){
			return []int{edge[0], edge[1]}
		}
	}
	return []int{}
}
