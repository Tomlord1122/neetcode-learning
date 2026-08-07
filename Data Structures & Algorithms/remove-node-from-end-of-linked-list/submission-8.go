/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
	dummy := &ListNode{Next:head}
	cur := dummy
	for cur.Next != nil{
		length++
		cur = cur.Next
	}
	cur = dummy
	for i := 0; i < length - n; i++{
		cur = cur.Next
	}
	// revmoe the next node
	cur.Next = cur.Next.Next
	return dummy.Next
}
