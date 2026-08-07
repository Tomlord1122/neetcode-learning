/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return valid(root, math.MinInt, math.MaxInt)
}

func valid (node *TreeNode, left, right int) bool{
	if node == nil{
		return true
	}
	val := node.Val
	if val <= left || val >= right{
		return false
	}
	return valid(node.Left, left, val) && valid(node.Right, val, right)
}

// dfs solution