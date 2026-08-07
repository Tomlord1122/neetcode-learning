/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	// HashMap -> oldToCopy [*Node]*Node
	oldToCopy := make(map[*Node]*Node)
	cur := head
	for cur != nil{
		cp := &Node{Val:cur.Val}
		oldToCopy[cur] = cp
		cur = cur.Next
	}
	cur = head
	for cur != nil{
		cp := oldToCopy[cur]
		cp.Next = oldToCopy[cur.Next]
		cp.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}
	return oldToCopy[head]
}
