/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	
	getLength := func(node *ListNode) int{
		res := 0
		for node != nil{
			res++
			node = node.Next
		}
		return res
	}

	m, n := getLength(headA), getLength(headB)
	l1, l2 := headA, headB
	
	if m < n{
		m, n = n, m
		l1, l2 = l2, l1
	}
	// Keep l1 as the longer one
	for m > n{
		l1 = l1.Next
		m--
	}
	for l1 != l2{
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1
}
