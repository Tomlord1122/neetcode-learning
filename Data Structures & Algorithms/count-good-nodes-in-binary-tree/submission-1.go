/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	if root == nil{
		return 0
	}
	var dfs func(node *TreeNode, curMax int) int
	dfs = func(node *TreeNode, curMax int) int{
		if node == nil{
			return 0
		}
		count := 0
		if node.Val >= curMax{
			count++
			curMax = node.Val
		}
		// recursive call left and right with new curMax
		count += dfs(node.Left, curMax)
		count += dfs(node.Right, curMax)
		return count
	}

	return dfs(root, root.Val)
}


// dfs -> pass the cur max to it