/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next{
		node := &Node{Val:cur.Val}
		oldToCopy[cur] = node
	}
	for cur := head; cur != nil ; cur = cur.Next{
		// process next and random
		node := oldToCopy[cur]
		node.Next = oldToCopy[cur.Next]
		node.Random = oldToCopy[cur.Random]
	}
	return oldToCopy[head]
}
