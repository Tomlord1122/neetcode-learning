func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make(map[int][]int)
	for i := 0; i < numCourses; i++{
		preMap[i] = []int{}
	}
	for _, course := range prerequisites{
		src, dst := course[0], course[1]
		preMap[src] = append(preMap[src], dst)
	}
	visited := make(map[int]bool)
	var dfs func(src int) bool
	dfs = func(src int) bool{
		if visited[src] == true{
			return false
		}
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
		// This can help reduce the time complexity
		preMap[src] = []int{} 
		return true
	}

	for c := 0 ; c < numCourses; c++{
		if !dfs(c){
			return false
		}
	}
	return true
}
