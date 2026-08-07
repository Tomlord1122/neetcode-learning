/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil{
		length++
		cur = cur.Next
	}
	cur = dummy
	for i := 0; i < length - n; i++{
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}


// 1 -> 2 -> 3 -> 4 => n = 2, length = 4
// dummy -> 1 -> 2 -> 3 -> 4  => length - n
// 1 -> 2 -> 3 -> 4
//      |
