/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	return dfs(root, math.MinInt, math.MaxInt)
}


func dfs(node *TreeNode, left int, right int) bool{
	if node == nil{
		return true
	}
	v := node.Val
	if v <= left || v >= right{
		return false
	}
	return dfs(node.Left, left, v) && dfs(node.Right, v, right)
}