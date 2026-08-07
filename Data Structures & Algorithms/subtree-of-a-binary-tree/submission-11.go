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
	if subRoot == nil{
		return true
	}
	if isSameTree(root, subRoot){
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(n1, n2 *TreeNode) bool{
	if n1 == nil && n2 == nil{
		return true
	}
	if n1 == nil || n2 == nil{
		return false
	}
	return n1.Val == n2.Val && isSameTree(n1.Left, n2.Left) && isSameTree(n1.Right, n2.Right)
}