type Graph struct {
	adjList map[int]map[int]bool // src -> dst 
}

func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int]map[int]bool),
	}
}

func (g *Graph) AddEdge(src, dst int) {
	if _, exist := g.adjList[src]; !exist{
		g.adjList[src] = make(map[int]bool)
	}
	if _, exist := g.adjList[dst]; !exist{
		g.adjList[dst] = make(map[int]bool)
	}
	g.adjList[src][dst] = true
}

func (g *Graph) RemoveEdge(src, dst int) bool {
	if _, exist := g.adjList[src]; !exist{
		return false
	}
	if _, exist := g.adjList[src][dst]; !exist{
		return false
	}
	delete(g.adjList[src], dst)
	return true
}

func (g *Graph) HasPath(src, dst int) bool {
	visited := make(map[int]bool)
	return g.dfs(src, dst, visited)
}

func (g *Graph) dfs(src, dst int, visited map[int]bool) bool{
	if src == dst{
		return true
	}
	visited[src] = true
	if neighbors, exist := g.adjList[src]; exist{
		for neighbor := range neighbors{
			if !visited[neighbor]{
				if g.dfs(neighbor, dst, visited){
					return true
				}
			}
		}
	}
	return false
}
