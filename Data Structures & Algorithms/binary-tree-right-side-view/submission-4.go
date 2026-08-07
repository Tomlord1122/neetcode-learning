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
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0{
		length := len(queue)
		for i := 0; i < length; i++{
			if i == length - 1{
				res = append(res, queue[0].Val)
			}
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil{
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil{
				queue = append(queue, pop.Right)
			}
		}
	}
	return res
}
