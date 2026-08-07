/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    for cur := head; cur != nil; cur = cur.Next{
        length++
    }
    step := length - n
    dummy := &ListNode{Next:head}
    cur := dummy
    for i := 0; i < step; i++{
        cur = cur.Next
    }
    cur.Next = cur.Next.Next
    return dummy.Next
}
