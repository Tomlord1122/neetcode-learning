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
	cur := slow
	var prev *ListNode
	for cur != nil{
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	l1, l2 := head, prev
	for l2.Next != nil{
		l1Next := l1.Next
		l2Next := l2.Next
		l1.Next = l2
		l2.Next = l1Next
		l1 = l1Next
		l2 = l2Next
	}
}



// 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6

// 0 -> 1 -> 2 -> 3 <- 4 <- 5 <- 6

// our result should be like this

// 0 -> 6 -> 1 -> 5 -> 2 -> 4 -> 3