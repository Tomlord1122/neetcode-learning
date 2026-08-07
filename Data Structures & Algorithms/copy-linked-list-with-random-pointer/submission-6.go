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
	cur := head
	// first round
	for cur != nil {
		oldToNew[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	// second round
	cur = head
	for cur != nil{
		oldToNew[cur].Next = oldToNew[cur.Next]
		oldToNew[cur].Random = oldToNew[cur.Random]
		cur = cur.Next
	}
	return oldToNew[head]
}


// hashMap -> key: head  value: newHead

// traverse this linked list twice
// O(n) -> 
// first round -> create hashMap
// second round -> update random pointer field