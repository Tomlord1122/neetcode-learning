func validTree(n int, edges [][]int) bool {
   neighborMap := make(map[int][]int) 
	for _, edge := range edges{
		src, dst := edge[0], edge[1]
		neighborMap[src] = append(neighborMap[src], dst)
		neighborMap[dst] = append(neighborMap[dst], src)
	}
	visit := make(map[int]bool)
	cycle := make(map[int]bool)

	var dfs func(i, parent int) bool
	dfs = func(i, parent int) bool{
		if cycle[i]{
			return false
		}
		if visit[i]{
			return true
		}
		cycle[i] = true
		for _, nei := range neighborMap[i]{
			if nei == parent{
				continue
			}
			if !dfs(nei, i){
				return false
			}
		}
		cycle[i] = false
		visit[i] = true
		return true
	}
	return dfs(0, -1) && len(visit) == n
}


// acyclic + connected => valid tree