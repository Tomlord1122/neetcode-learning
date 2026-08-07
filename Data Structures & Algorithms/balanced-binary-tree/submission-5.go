/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    res := dfs(root)
    return res.balanced
}

func dfs(node *TreeNode) result{
    if node == nil{
        return result{
            balanced: true,
            height: 0,
        }
    }
    left := dfs(node.Left)
    right := dfs(node.Right)
    return result{
        balanced: left.balanced && right.balanced && abs(left.height - right.height) < 2,
        height: 1 + max(left.height, right.height),
    }
}


// check left subtree, right subtree is balanced?
// We need to get the length of left/right subtree
type result struct{
    balanced bool
    height int
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
}