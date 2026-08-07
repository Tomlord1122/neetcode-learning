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
	// first pass
	cur := head
	for cur != nil{
		cp := &Node{Val:cur.Val}
		oldToCopy[cur] = cp
		cur = cur.Next
	}

	// Second pass
	// Deal with the random pointer
	cur = head
	for cur != nil{
		cp := oldToCopy[cur]
		cp.Next = oldToCopy[cur.Next]
		cp.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}
	return oldToCopy[head]
}
// first pass
// create the deep copy of node 
// create a hashMap to map old to copy

// second pass