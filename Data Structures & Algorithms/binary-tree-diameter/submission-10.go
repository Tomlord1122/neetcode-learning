/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil{
		return 0
	}
	lHeight := getHeight(root.Left)
	rHeight := getHeight(root.Right)
	return max(lHeight+rHeight, diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
}

func getHeight(root *TreeNode) int{
	if root == nil{
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}
