/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    res := 0
	var dfs func(node*TreeNode, curMax int)
	dfs = func(node *TreeNode, curMax int){
		if node == nil{
			return
		}
		if node.Val >= curMax{
			res++
			curMax = node.Val
		}
		dfs(node.Left, curMax)
		dfs(node.Right, curMax)
	}

	dfs(root, root.Val)
	return res
}
