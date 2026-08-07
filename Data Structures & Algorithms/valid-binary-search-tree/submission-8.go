/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    var helper func(node *TreeNode, left int, right int) bool
	helper = func(node *TreeNode, left int, right int) bool{
		if node == nil{
			return true
		}
		if node.Val <= left || node.Val >= right{
			return false
		}
		return helper(node.Left, left, node.Val) && helper(node.Right, node.Val, right)
	}
	return helper(root, math.MinInt, math.MaxInt)
}
