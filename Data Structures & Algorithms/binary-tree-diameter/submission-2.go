/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    diameter := 0

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int{
        if node == nil{
            return 0
        }
        leftHeight := dfs(node.Left)
        rightHeight := dfs(node.Right)
        diameter = max(diameter, leftHeight + rightHeight)
        return 1 + max(leftHeight, rightHeight)
    }
    dfs(root)
    return diameter
}
