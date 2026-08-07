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
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0{
		length := len(queue)
		curLevel := []int{}
		for i := 0; i < length; i++{
			cur := queue[0]
			queue = queue[1:]
			curLevel = append(curLevel, cur.Val)
			if cur.Left != nil{
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil{
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, curLevel)
	}
	return res
}
