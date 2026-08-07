/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
    res := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		cur := []int{}
		for i := 0; i < length; i++{
			pop := queue[0]
			queue = queue[1:]
			if pop.Left != nil{
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil{
				queue = append(queue, pop.Right)
			}
			cur = append(cur, pop.Val)
		}
		res = append(res, cur)
	}
	return res
}
