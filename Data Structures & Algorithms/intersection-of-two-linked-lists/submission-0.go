/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	getLength := func(head *ListNode) int{
		length := 0
		for head != nil{
			length++
			head = head.Next
		}
		return length
	}

	m, n := getLength(headA), getLength(headB)
	l1, l2 := headA, headB
	if m < n{
		m, n = n, m
		l1, l2 = l2, l1
	}

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