/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    if root == nil{
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int{
		if node == nil{
			return 0
		}

		left, right := dfs(node.Left), dfs(node.Right)
		return 1 + max(left, right)
	}
	return dfs(root)
}
