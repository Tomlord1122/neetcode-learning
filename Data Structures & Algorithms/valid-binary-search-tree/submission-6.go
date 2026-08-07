/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    
	var dfs func(node *TreeNode, l int, r int) bool
	dfs = func(node *TreeNode, l int, r int) bool{
		if node == nil{
			return true
		}
		if node.Val <= l || node.Val >= r{
			return false
		}
		return dfs(node.Left, l, node.Val) && dfs(node.Right, node.Val, r)
	}
	return dfs(root, math.MinInt, math.MaxInt)
}
