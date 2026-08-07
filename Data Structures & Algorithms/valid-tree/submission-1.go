func validTree(n int, edges [][]int) bool {
    // create the preMap to represent the graph (adjcency list)
	preMap := make(map[int][]int)
	for _, edge := range edges{
		src, dst := edge[0], edge[1]
		// the graph is a undirected graph
		// so we need to append both of the direction
		preMap[src] = append(preMap[src], dst)
		preMap[dst] = append(preMap[dst], src)
	}

	// Then we use a dfs function to check the cycle in each dfs call
	// and maintain a visit set to check if the graph is connected
	cycle := make(map[int]bool)
	visit := make(map[int]bool)
	var dfs func(node int, parent int) bool 
	dfs = func(node int, parent int) bool{
		if cycle[node]{
			return false
		}
		if visit[node]{
			return true
		}
		cycle[node] = true
		for _, nei := range preMap[node]{
			if nei == parent{
				continue
			}
			if !dfs(nei, node){
				return false
			}
		}
		cycle[node] = false
		visit[node] = true
		return true
	}

	return dfs(0, -1) && len(visit) == n
}


// what's a valid validTree
// This graph does not have a cycle and it's connected.