/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    res := []int{}
	var ldr func(node *TreeNode)
	ldr = func(node *TreeNode){
		if node == nil{
			return 
		}
		ldr(node.Left)
		res = append(res, node.Val)
		ldr(node.Right)
	}
	ldr(root)
	return res[k-1]
}
