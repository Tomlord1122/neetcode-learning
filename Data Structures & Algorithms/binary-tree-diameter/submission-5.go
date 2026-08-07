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
	if root == nil{
		return 0
	}

	var treeHeight func(node *TreeNode) int
	treeHeight = func(node *TreeNode) int{
		if node == nil{
			return 0
		}
		leftH := treeHeight(node.Left)
		rightH := treeHeight(node.Right)
		diameter = max(diameter, leftH + rightH)
		return 1 + max(leftH, rightH)
	}

	treeHeight(root)
	return diameter
}
