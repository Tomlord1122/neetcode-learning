/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs (node *TreeNode, curMax int) int{
	if node == nil{
		return 0
	}
	count := 0
	if node.Val >= curMax{
		count++
		curMax = node.Val
	}
	count += dfs(node.Left, curMax) + dfs(node.Right, curMax)
	return count
}