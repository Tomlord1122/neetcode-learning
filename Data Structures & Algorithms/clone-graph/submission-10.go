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
	oldToNew[node] = &Node{Val: node.Val}
	queue := []*Node{node}
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++{
			cur := queue[0]
			queue = queue[1:]
			for _, nei := range cur.Neighbors{
				if _, exist := oldToNew[nei]; !exist{
					oldToNew[nei] = &Node{
						Val: nei.Val,
					}
					// append to queue because this node occur first time
					queue = append(queue, nei)
				}
				oldToNew[cur].Neighbors = append(
					oldToNew[cur].Neighbors,
					oldToNew[nei],
				)
			}
		}
	}
	return oldToNew[node]
}
