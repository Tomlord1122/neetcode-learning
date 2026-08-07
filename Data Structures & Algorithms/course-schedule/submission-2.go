func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++{
		preMap[i] = []int{}
	}
	for _, prereq := range prerequisites{
		src, dst := prereq[0], prereq[1]
		preMap[src] = append(preMap[src], dst)
	}

	visited := make(map[int]bool)

	var dfs func(int) bool
	dfs = func(src int) bool{
		// What's the base case?
		// Found the cycle
		if visited[src]{
			return false
		}
		// The current src does not have any dst
		if preMap[src] == nil{
			return true
		}
		visited[src] = true
		for _, dst := range preMap[src]{
			if !dfs(dst){
				return false
			}
		}
		visited[src] = false
		preMap[src] = []int{}
		return true
	}

	for i := 0; i < numCourses; i++{
		if !dfs(i){
			return false
		}
	}
	return true
}
