/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil{
		v1, v2 := list1.Val, list2.Val
		if v1 < v2{
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil{
		cur.Next = list1
		cur = cur.Next
	}
	if list2 != nil{
		cur.Next = list2
		cur = cur.Next
	}
	return dummy.Next
}
