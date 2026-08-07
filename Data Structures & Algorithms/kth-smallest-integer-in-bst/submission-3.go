/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    // preorder traversal -> get sorted array -> return arr[k-1]
	arr := []int{}
	var ldr func(node *TreeNode)
	ldr = func(node *TreeNode){
		if node == nil{
			return
		}
		ldr(node.Left)
		arr = append(arr, node.Val)
		ldr(node.Right)
	}
	ldr(root)
	return arr[k-1]
}
