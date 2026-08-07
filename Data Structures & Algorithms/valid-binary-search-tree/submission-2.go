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

func valid(node *TreeNode, leftCond int, rightCond int) bool{
	if node == nil{
		return true
	}
	if node.Val <= leftCond || node.Val >= rightCond{
		return false
	}
	return valid(node.Left, leftCond, node.Val) && valid(node.Right, node.Val, rightCond)
}