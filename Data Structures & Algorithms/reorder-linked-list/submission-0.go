/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    if head == nil || head.Next == nil{
        return
    }
    slow, fast := head, head
    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    // Reverse the second half
    cur := slow.Next
    // Cut first half at slow
    slow.Next = nil
    var prev *ListNode
    for cur != nil {
        nxt := cur.Next
        cur.Next = prev
        prev = cur
        cur = nxt
    }
    second := prev
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
