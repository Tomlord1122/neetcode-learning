/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	cur := slow
	var prev *ListNode
	// reverse the second half
	for cur != nil{
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	head1, head2 := head, prev
	for head2 != nil{
		if head1.Val != head2.Val{
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}
	return true
}
