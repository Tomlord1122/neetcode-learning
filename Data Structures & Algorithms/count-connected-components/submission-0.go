func countComponents(n int, edges [][]int) int {
    visit := make(map[int]bool)
	res := 0

	neighborMap := make(map[int][]int)
	for _, edge := range edges{
		src, dst := edge[0], edge[1]
		neighborMap[src] = append(neighborMap[src], dst)
		neighborMap[dst] = append(neighborMap[dst], src)
	}

	var dfs func(i int)
	dfs = func(i int){
		if visit[i]{
			return 
		}
		visit[i] = true
		for _, nei := range neighborMap[i]{
			dfs(nei)
		}
	}

	for i := 0; i < n; i++{
		if !visit[i]{
			dfs(i)
			res++
		}
	}
	return res
}
