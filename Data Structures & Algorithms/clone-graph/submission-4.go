/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil{
		return nil
	}
    oldToNew := make(map[*Node]*Node)
	oldToNew[node] = &Node{Val:node.Val, Neighbors: []*Node{}}
	queue := []*Node{node}
	for len(queue) > 0{
		// iterate the node's Neighbors
		pop := queue[0]
		queue = queue[1:]
		for _, nei := range pop.Neighbors{
			if _, exist := oldToNew[nei]; !exist{
				oldToNew[nei] = &Node{Val:nei.Val, Neighbors: []*Node{}}
				queue = append(queue, nei)
			}
			// update oldToNew[node]'s Neighbors
			oldToNew[pop].Neighbors = append(oldToNew[pop].Neighbors, oldToNew[nei])
		}
	}
	return oldToNew[node]
}
