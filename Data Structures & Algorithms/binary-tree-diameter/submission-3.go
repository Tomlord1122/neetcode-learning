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
    dfs = func(node *TreeNode) int {
        if node == nil{
            return 0
        }
        leftH := dfs(node.Left)
        rightH := dfs(node.Right)
        diameter = max(diameter, leftH + rightH)
        return 1 + max(leftH, rightH)
    }
    dfs(root)
    return diameter
}


