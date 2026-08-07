func canFinish(numCourses int, prerequisites [][]int) bool {
	// Init the graph first
	preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++{
		preMap[i] = []int{}
	}
	for _, prereq := range prerequisites{
		src, dst := prereq[0], prereq[1]
		preMap[src] = append(preMap[src], dst)
	}
	visited := make(map[int]bool)
	var dfs func(node int) bool
	dfs = func(node int) bool{
		// base case
		if visited[node]{
			return false
		}
		if preMap[node] == nil{
			return true
		}
		visited[node] = true
		for _, nei := range preMap[node]{
			if !dfs(nei){
				return false
			}
		}
		preMap[node] = []int{}
		visited[node] = false
		return true
	}
	// iterate all node and check if there is a cycle
	for i := 0; i < numCourses; i++{
		if !dfs(i){
			return false
		}
	}
	return true
}
