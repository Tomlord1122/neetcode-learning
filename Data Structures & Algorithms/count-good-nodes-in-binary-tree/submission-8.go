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
	res := 0

	var dfs func(node *TreeNode, curMax int)
	dfs = func(node *TreeNode, curMax int){
		if node == nil{
			return 
		}
		if node.Val >= curMax{
			curMax = node.Val
			res++
		}
		if node.Left != nil{
			dfs(node.Left, curMax)
		}
		if node.Right != nil{
			dfs(node.Right, curMax)
		}
	}

	dfs(root, root.Val)
	return res
}
