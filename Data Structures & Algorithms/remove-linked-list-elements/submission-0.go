/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	prev, cur := dummy, head
	for cur != nil{
		nxt := cur.Next
		if cur.Val == val{
			prev.Next = nxt
		} else {
			prev = cur
		}
		cur = nxt
	}
	return dummy.Next
}
