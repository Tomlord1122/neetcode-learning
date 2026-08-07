func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := make(map[int][]int)
    for _, prereq := range prerequisites{
        src, dst := prereq[0], prereq[1]
        preMap[src] = append(preMap[src], dst)
    }

    visit, cycle := make(map[int]bool), make(map[int]bool)
    res := []int{}
    var dfs func(course int) bool
    dfs = func(course int) bool{
        // base case
        if cycle[course]{
            return false
        }
        if visit[course]{
            return true
        }
        cycle[course] = true
        for _, nei := range preMap[course]{
            if !dfs(nei){
                return false
            }
        }
        cycle[course] = false
        // append the current course into res array
        res = append(res, course)
        visit[course] = true
        return true
    }

    for i := 0; i < numCourses; i++{
        if !dfs(i){
            return []int{}
        }
    }
    return res
}
