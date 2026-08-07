/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int{
		if node == nil{
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		diameter = max(diameter, left + right)
		return 1 + max(left, right)
	}
	helper(root)
	return diameter
}
