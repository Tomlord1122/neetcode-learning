/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    arr := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode){
		if node == nil{
			return
		}
		// LDR
		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return arr[k-1]
}
