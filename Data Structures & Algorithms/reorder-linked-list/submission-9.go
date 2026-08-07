/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	second := slow.Next
	slow.Next = nil
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
		fNext := first.Next
		sNext := second.Next
		first.Next = second
		second.Next = fNext
		first = fNext
		second = sNext
	}
}
//                |
// 1 -> 2 -> 3 -> 4 -> 5 -> 6
// 1 -> 2 -> 3 -> 4 (first)
// 6 -> 5 (second)
//                |
// 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7
// 1 -> 2 -> 3 -> 4
// 7 -> 6 -> 5