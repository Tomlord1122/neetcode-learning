func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
	for _, prereq := range prerequisites{
		src, dst := prereq[0], prereq[1]
		preMap[src] = append(preMap[src], dst)
	}
	// Now we have the graph representation (adjacency list)
	var dfs func(src int) bool
	visited := make(map[int]bool)
	dfs = func(src int) bool{
		if visited[src]{
			return false
		}
		if preMap[src] == nil{
			return true
		}
		visited[src] = true
		for _, nei := range preMap[src]{
			if !dfs(nei){
				return false
			}
		}
		visited[src] = false
		preMap[src] = nil
		return true
	}

	for i := 0; i < numCourses; i++{
		if !dfs(i){
			return false
		}
	}
	return true
}
