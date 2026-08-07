/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, left int, right int) bool
	dfs = func(node *TreeNode, left int, right int) bool{
		if node == nil{
			return true
		}
		if node.Val <= left || node.Val >= right{
			return false
		}
		return dfs(node.Left, left, node.Val) && dfs(node.Right, node.Val, right)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
