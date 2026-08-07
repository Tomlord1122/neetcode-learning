/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    if root == nil{
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil{
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil{
				queue = append(queue, cur.Right)
			}
			if i == length-1{
				res = append(res, cur.Val)
			}
		}
	}
	
	return res
}
