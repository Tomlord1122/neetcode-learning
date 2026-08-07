/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	getLength := func(node *ListNode) int{
		if node == nil{
			return 0
		}
		res := 0
		for node != nil{
			res++
			node = node.Next
		}
		return res
	}	

	lengthA, lengthB := getLength(headA), getLength(headB)
	if lengthA < lengthB{
		lengthA, lengthB = lengthB, lengthA
		headA, headB = headB, headA
	}
	for i := 0; i < lengthA - lengthB; i++{
		headA = headA.Next
	}

	for headA != headB{
		headA = headA.Next
		headB = headB.Next
	}

	return headA
}
