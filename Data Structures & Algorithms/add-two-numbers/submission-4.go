/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	car := 0
	for l1 != nil || l2 != nil || car != 0{
		l1Val, l2Val := 0, 0
		if l1 != nil{
			l1Val = l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			l2Val = l2.Val
			l2 = l2.Next
		}
		res := l1Val + l2Val + car
		car = res / 10
		cur.Next = &ListNode{Val:res%10}
		cur = cur.Next
	}

	return dummy.Next
}