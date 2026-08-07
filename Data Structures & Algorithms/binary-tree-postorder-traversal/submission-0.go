/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    res := []int{}

	var helper func(node *TreeNode)
	helper = func(node *TreeNode){
		if node == nil{
			return 
		}
		helper(node.Left)
		helper(node.Right)
		res = append(res, node.Val)
		return 
	}
	helper(root)
	return res
}