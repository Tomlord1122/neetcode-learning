/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next:head}
    length := 0
    for cur := head; cur != nil; cur = cur.Next{
        length++
    }
    step := length - n
    cur := dummy
    for i := 0; i < step; i++{
        cur = cur.Next
    }
    // right now cur is located before target node 
    cur.Next = cur.Next.Next
    return dummy.Next
}
