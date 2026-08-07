/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	oldToNew := make(map[*Node]*Node)
	// First pass
	cur := head
	for cur != nil{
		oldToNew[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil{
		oldToNew[cur].Next = oldToNew[cur.Next]
		oldToNew[cur].Random = oldToNew[cur.Random]
		cur = cur.Next
	}
	return oldToNew[head]
}

