/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil{
        return false
    }
    if subRoot == nil{
        return true
    }
    if isSameTree(root, subRoot){
        return true
    }
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}


func isSameTree(root *TreeNode, subRoot *TreeNode) bool{
    if root == nil && subRoot == nil{
        return true
    }
    if root != nil && subRoot !=nil && root.Val == subRoot.Val{
        return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
    }
    return false
}