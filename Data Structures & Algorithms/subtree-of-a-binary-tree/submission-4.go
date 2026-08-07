/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }
    if subRoot == nil{
        return true
    }
    if sameTree(root, subRoot){
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(node *TreeNode, subRoot *TreeNode) bool{
    if node == nil && subRoot == nil{
        return true
    }
    if node != nil && subRoot != nil && node.Val == subRoot.Val{
        return sameTree(node.Left, subRoot.Left) && sameTree(node.Right, subRoot.Right)
    }
    return false
}