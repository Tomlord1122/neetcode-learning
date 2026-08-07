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

    dummy := &ListNode{Next: head}
    prev := dummy
    steps := length - n
    for i := 0; i < steps; i++{
        prev = prev.Next
    }

    // if prev.Next != nil{
        prev.Next = prev.Next.Next
    // }
    return dummy.Next
}
