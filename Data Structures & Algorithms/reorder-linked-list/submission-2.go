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
    // reverse the second half linked list
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


// 0, 1, 2, 3, 4, 5
// 0, 1, 2,
// 5, 4, 3

// 0, 1, 2, 3, 4, 5, 6
// 0, 1, 2
// 6, 5, 4, 3