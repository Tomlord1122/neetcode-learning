/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }

    var prev *ListNode
    cur := head
    for cur != nil{
        nxt := cur.Next
        cur.Next = prev
        prev = cur
        cur = nxt
    }
    return prev
}
