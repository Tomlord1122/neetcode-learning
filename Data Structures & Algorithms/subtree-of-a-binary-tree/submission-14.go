/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil{
		return false
	}
	if subRoot == nil || isSame(root,subRoot){
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(root *TreeNode, node *TreeNode) bool{
	if root == nil && node == nil{
		return true
	}
	if root == nil || node == nil{
		return false
	}
	return root.Val == node.Val && isSame(root.Left, node.Left) && isSame(root.Right, node.Right) 
}