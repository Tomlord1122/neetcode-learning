/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	// reverse second half
	// reorder it
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	// reverse second half
	var prev *ListNode
	for second != nil{
		nxt := second.Next
		second.Next = prev
		prev = second
		second = nxt
	}
	second = prev
	first := head
	for second != nil{
		sNext := second.Next
		fNext := first.Next
		first.Next = second
		second.Next = fNext
		first = fNext
		second = sNext
	}
}
