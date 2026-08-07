/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    if root == nil{
		return true
	}
	var dfs func(node *TreeNode, leftBond int, rightBond int) bool
	dfs = func(node *TreeNode, leftBond int, rightBond int) bool{
		if node == nil{
			return true
		}
		if node.Val <= leftBond || node.Val >= rightBond{
			return false
		}
		return dfs(node.Left, leftBond, node.Val) && dfs(node.Right, node.Val, rightBond)
	}

	// call the dfs function at root
	// this will check all entire tree
	return dfs(root, math.MinInt, math.MaxInt)
}
